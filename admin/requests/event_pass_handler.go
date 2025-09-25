package requests

import (
	"net/http"
	"text/template"
    "time"

	"test2/utils"
	"github.com/labstack/echo/v4"
)

type EventPass struct {
	ID        string
	GroupID   string
	Code      string
	StartDate string
	EndDate   string
}

func EventPassHandler(c echo.Context) error {
	db, err := utils.OpenDB()
	if err != nil {
		return c.String(http.StatusInternalServerError, "DB 연결 오류")
	}
	defer db.Close()

	rows, err := db.Query("SELECT ID, GROUP_ID, CODE, START_DATE, END_DATE FROM event_pass ORDER BY START_DATE DESC")
	if err != nil {
		return c.String(http.StatusInternalServerError, "이벤트 목록 조회 오류")
	}
	defer rows.Close()

	var events []EventPass
	for rows.Next() {
		var event EventPass
        var startDate, endDate int64
		if err := rows.Scan(&event.ID, &event.GroupID, &event.Code, &startDate, &endDate); err != nil {
			return c.String(http.StatusInternalServerError, "이벤트 목록 스캔 오류")
		}
        event.StartDate = time.Unix(startDate, 0).Format("2006-01-02 15:04:05")
        event.EndDate = time.Unix(endDate, 0).Format("2006-01-02 15:04:05")
		events = append(events, event)
	}

	tmpl, err := template.ParseFiles("pages/event_pass.html")
	if err != nil {
		return c.String(http.StatusInternalServerError, "템플릿 로드 오류")
	}

	return tmpl.Execute(c.Response().Writer, map[string]interface{}{
		"Events": events,
	})
}
