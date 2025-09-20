package request

import (
    "database/sql"
    "log"
    "net/http"
    "test1/utils"

    "github.com/labstack/echo/v4"
)

func LoginActionHandler(c echo.Context) error {
    username := c.FormValue("username")
    password := c.FormValue("password")

    db, err := utils.GetDB()
    if err != nil {
        log.Printf("Error opening database connection: %v", err)
        return c.Redirect(http.StatusFound, "/login?error=2") // DB connection error
    }
    defer db.Close()

    var dbPassword string
    err = db.QueryRow("SELECT password FROM admin_auth WHERE user_name = ?", username).Scan(&dbPassword)

    if err != nil {
        if err == sql.ErrNoRows {
            // User not found
            log.Printf("Login failed for user: %s, user not found", username)
            return c.Redirect(http.StatusFound, "/login?error=1")
        }
        // Other query error
        log.Printf("Error querying database: %v", err)
        return c.Redirect(http.StatusFound, "/login?error=2")
    }

    // In a real application, you should use bcrypt or another hashing algorithm.
    // Storing plain text passwords is a major security risk.
    if password == dbPassword {
        // Login successful
        log.Printf("Login successful for user: %s", username)
        return c.Redirect(http.StatusFound, "/admin_list")
    }

    // Password mismatch
    log.Printf("Login failed for user: %s, password mismatch", username)
    return c.Redirect(http.StatusFound, "/login?error=1")
}
