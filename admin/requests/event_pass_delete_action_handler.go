package requests

import (
	"net/http"
	"test2/utils"

	"github.com/labstack/echo/v4"
)

func EventPassDeleteActionHandler(c echo.Context) error {
	id := c.QueryParam("id")

	db, err := utils.OpenDB()
	if err != nil {
		return c.String(http.StatusInternalServerError, "DB 연결 오류")
	}
	defer db.Close()

	_, err = db.Exec("DELETE FROM event_pass WHERE ID = ?", id)
	if err != nil {
		return c.String(http.StatusInternalServerError, "이벤트 삭제 오류")
	}

	return c.Redirect(http.StatusFound, "/event/pass")
}
