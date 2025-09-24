package requests

import (
	"database/sql"
	"net/http"

	"test2/utils"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

func LoginActionHandler(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	db, err := utils.OpenDB()
	if err != nil {
		c.Logger().Errorf("Error opening database: %v", err)
		return c.String(http.StatusInternalServerError, "데이터베이스 연결 오류")
	}
	defer db.Close()

	var hashedPassword string
	err = db.QueryRow("SELECT password FROM admin_auth WHERE user_name = ?", username).Scan(&hashedPassword)
	if err != nil {
		if err == sql.ErrNoRows {
			return c.HTML(http.StatusOK, "<script>alert('로그인 실패하였습니다.'); window.location.href='/login';</script>")
		}
		c.Logger().Errorf("Error querying user: %v", err)
		return c.String(http.StatusInternalServerError, "사용자 조회 오류")
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return c.HTML(http.StatusOK, "<script>alert('로그인 실패하였습니다.'); window.location.href='/login';</script>")
	}

	// Successful login, store in session and redirect to admin_list
	sess, _ := session.Get("session", c)
	sess.Values["authenticated"] = true
	sess.Save(c.Request(), c.Response())
	return c.Redirect(http.StatusFound, "/admin_list")
}
