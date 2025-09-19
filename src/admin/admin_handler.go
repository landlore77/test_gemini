package admin

import (
	"database/sql"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	_ "github.com/go-sql-driver/mysql" // MySQL driver
	"test1/model" // Corrected module name
)

const (
	adminsPerPage = 10
)

// AdminListHandler handles the /admin_list page
func AdminListHandler(c echo.Context) error {
	// TODO: Get database connection from context or global variable
	// For now, a placeholder connection
	db, err := sql.Open("mysql", "test1:test1@tcp(127.0.0.1:3306)/test_admin")
	if err != nil {
		log.Printf("Error opening database connection: %v", err)
		return c.String(http.StatusInternalServerError, "Error connecting to database")
	}
	defer db.Close()

	pageStr := c.QueryParam("page")
	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}

	offset := (page - 1) * adminsPerPage

	// Fetch total count for pagination
	var totalAdmins int
	err = db.QueryRow("SELECT COUNT(*) FROM admin_auth").Scan(&totalAdmins)
	if err != nil {
		log.Printf("Error counting admins: %v", err)
		return c.String(http.StatusInternalServerError, "Error fetching admin count")
	}

	totalPages := (totalAdmins + adminsPerPage - 1) / adminsPerPage

	// Fetch admins for the current page
	rows, err := db.Query("SELECT id, group_id, user_name, description, last_ip, last_login FROM admin_auth ORDER BY create_time DESC LIMIT ? OFFSET ?", adminsPerPage, offset)
	if err != nil {
		log.Printf("Error fetching admins: %v", err)
		return c.String(http.StatusInternalServerError, "Error fetching admin data")
	}
	defer rows.Close()

	var admins []model.AdminDisplayData
	for rows.Next() {
		var admin model.Admin
		err := rows.Scan(&admin.ID, &admin.GroupID, &admin.UserName, &admin.Description, &admin.LastIP, &admin.LastLogin)
		if err != nil {
			log.Printf("Error scanning admin row: %v", err)
			return c.String(http.StatusInternalServerError, "Error processing admin data")
		}

		// Format last_login timestamp
		formattedLastLogin := time.Unix(int64(admin.LastLogin), 0).Format("2006-01-02 15:04:05")

		admins = append(admins, model.AdminDisplayData{
			ID:                 admin.ID,
			GroupID:            admin.GroupID,
			UserName:           admin.UserName,
			Description:        admin.Description,
			LastIP:             admin.LastIP,
			FormattedLastLogin: formattedLastLogin,
		})
	}

	// Generate pagination links
	var pages []int
	for i := 1; i <= totalPages; i++ {
		pages = append(pages, i)
	}

	data := model.AdminListPageData{
		Admins:      admins,
		CurrentPage: page,
		TotalPages:  totalPages,
		PrevPage:    page - 1,
		NextPage:    page + 1,
		Pages:       pages,
	}

	// Render the template using Echo's renderer
	return c.Render(http.StatusOK, "admin_list.html", data)
}