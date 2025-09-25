package requests

import (
	"net/http"
	"test2/utils"
	"time"

	"github.com/labstack/echo/v4"
)

func EventPassAddActionHandler(c echo.Context) error {
	groupID := c.FormValue("group_id")
	code := c.FormValue("code")
	startDateStr := c.FormValue("start_date")
	endDateStr := c.FormValue("end_date")

	layout := "2006-01-02 15:04:05"
	startTime, err := time.Parse(layout, startDateStr)
	if err != nil {
		return c.String(http.StatusBadRequest, "잘못된 시작일시 형식입니다.")
	}
	endTime, err := time.Parse(layout, endDateStr)
	if err != nil {
		return c.String(http.StatusBadRequest, "잘못된 종료일시 형식입니다.")
	}

	db, err := utils.OpenDB()
	if err != nil {
		return c.String(http.StatusInternalServerError, "DB 연결 오류")
	}
	defer db.Close()

	id := utils.GenerateID()

	_, err = db.Exec("INSERT INTO event_pass (ID, GROUP_ID, CODE, START_DATE, END_DATE) VALUES (?, ?, ?, ?, ?)",
		id, groupID, code, startTime.Unix(), endTime.Unix())
	if err != nil {
		return c.String(http.StatusInternalServerError, "이벤트 추가 오류")
	}

	return c.Redirect(http.StatusFound, "/event/pass")
}
