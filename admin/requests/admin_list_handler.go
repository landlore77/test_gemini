package requests

import (
	"net/http"
	"strconv"
	"test2/utils"
	"time"

	"github.com/labstack/echo/v4"
	_ "github.com/go-sql-driver/mysql"
)

type Admin struct {
	ID          int
	GroupID     int
	UserName    string
	Description string
	LastLogin   string // Formatted time string
	LastIP      string
}

type AdminListPageData struct {
	Admins     []Admin
	CurrentPage int
	TotalPages  int
	PrevPage    int
	NextPage    int
}

func AdminListHandler(c echo.Context) error {
	db, err := utils.OpenDB()
	if err != nil {
		c.Logger().Errorf("Error opening database: %v", err)
		return c.String(http.StatusInternalServerError, "Error connecting to database")
	}
	defer db.Close()

	pageStr := c.QueryParam("page")
	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}

	limit := 10
	offset := (page - 1) * limit

	// Get total count for pagination
	var totalAdmins int
	err = db.QueryRow("SELECT COUNT(*) FROM admin_auth").Scan(&totalAdmins)
	if err != nil {
		c.Logger().Errorf("Error fetching admin count: %v", err)
		return c.String(http.StatusInternalServerError, "Error fetching admin count")
	}

	totalPages := (totalAdmins + limit - 1) / limit
	if totalPages == 0 {
		totalPages = 1
	}

	rows, err := db.Query("SELECT id, group_id, user_name, description, last_login, last_ip FROM admin_auth ORDER BY create_time DESC LIMIT ? OFFSET ?", limit, offset)
	if err != nil {
		c.Logger().Errorf("Error querying admins: %v", err)
		return c.String(http.StatusInternalServerError, "Error querying admins")
	}
	defer rows.Close()

	admins := []Admin{}
	for rows.Next() {
		var admin Admin
		var lastLoginInt int64
		err := rows.Scan(&admin.ID, &admin.GroupID, &admin.UserName, &admin.Description, &lastLoginInt, &admin.LastIP)
		if err != nil {
			c.Logger().Errorf("Error scanning admin row: %v", err)
			return c.String(http.StatusInternalServerError, "Error scanning admin row")
		}
		admin.LastLogin = time.Unix(lastLoginInt, 0).Format("2006-01-02 15:04:05")
		admins = append(admins, admin)
	}

	data := AdminListPageData{
		Admins:      admins,
		CurrentPage: page,
		TotalPages:  totalPages,
		PrevPage:    page - 1,
		NextPage:    page + 1,
	}

	return c.Render(http.StatusOK, "admin_list.html", data)
}
