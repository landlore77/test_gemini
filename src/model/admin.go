package model

import (
	"time"
)

// Admin represents a row in the admin_auth table
type Admin struct {
	ID          int       `db:"id"`
	GroupID     int       `db:"group_id"`
	UserName    string    `db:"user_name"`
	Description string    `db:"description"`
	Password    string    `db:"password"`
	CreateTime  time.Time `db:"create_time"`
	LastIP      string    `db:"last_ip"`
	LastLogin   int       `db:"last_login"` // Unix timestamp
}

// AdminDisplayData is a struct for displaying admin data in the template
type AdminDisplayData struct {
	ID                 int
	GroupID            int
	UserName           string
	Description        string
	LastIP             string
	FormattedLastLogin string
}

// AdminListPageData is the data structure passed to the admin_list.html template
type AdminListPageData struct {
	Admins      []AdminDisplayData
	CurrentPage int
	TotalPages  int
	PrevPage    int
	NextPage    int
	Pages       []int // For pagination links
}