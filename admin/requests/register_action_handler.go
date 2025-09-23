package requests

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

func RegisterActionHandler(c echo.Context) error {
	userName := c.FormValue("user_name")
	password := c.FormValue("password")

	if len(userName) < 4 {
		return c.String(http.StatusBadRequest, "유저 이름이 너무 짧습니다.")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		c.Logger().Errorf("Error hashing password: %v", err)
		return c.String(http.StatusInternalServerError, "비밀번호 해싱 오류")
	}

	clientIP := c.RealIP()

	db, err := sql.Open("mysql", "test1:test1@tcp(127.0.0.1:3306)/test_admin")
	if err != nil {
		c.Logger().Errorf("Error opening database: %v", err)
		return c.String(http.StatusInternalServerError, "데이터베이스 연결 오류")
	}
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO admin_auth (group_id, user_name, description, password, last_ip, last_login) VALUES (?, ?, ?, ?, ?, ?)")
	if err != nil {
		c.Logger().Errorf("Error preparing statement: %v", err)
		return c.String(http.StatusInternalServerError, "데이터베이스 준비 오류")
	}
	defer stmt.Close()

	_, err = stmt.Exec(-1, userName, "신규유저", string(hashedPassword), clientIP, 0)
	if err != nil {
		c.Logger().Errorf("Error inserting admin: %v", err)
		return c.String(http.StatusInternalServerError, fmt.Sprintf("관리자 등록 오류: %v", err))
	}

	return c.Redirect(http.StatusSeeOther, "/admin_list")
}
