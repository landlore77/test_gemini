package request

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	_ "github.com/go-sql-driver/mysql" // MySQL driver
)

// RegisterRequest represents the request body for registration
type RegisterRequest struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

// RegisterActionHandler handles the /actions/register endpoint
func RegisterActionHandler(c echo.Context) error {
	req := new(RegisterRequest)
	if err := c.Bind(req); err != nil {
		log.Printf("Error binding request: %v", err)
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid request body"})
	}

	if req.UserName == "" || req.Password == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "User name and password cannot be empty"})
	}

	// TODO: Get database connection from context or global variable
	db, err := sql.Open("mysql", "test1:test1@tcp(127.0.0.1:3306)/test_admin")
	if err != nil {
		log.Printf("Error opening database connection: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Error connecting to database"})
	}
	defer db.Close()

	// Get client IP
	lastIP := c.RealIP()
	if lastIP == "::1" || lastIP == "127.0.0.1" { // Handle localhost for testing
		lastIP = "127.0.0.1"
	}

	// Insert into admin_auth table
	_, err = db.Exec(
		"INSERT INTO admin_auth (group_id, user_name, description, password, create_time, last_ip, last_login) VALUES (?, ?, ?, ?, ?, ?, ?)",
		-1, req.UserName, "신규유저", req.Password, time.Now(), lastIP, time.Now().Unix(),
	)
	if err != nil {
		log.Printf("Error inserting admin: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Error registering user"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "User registered successfully"})
}