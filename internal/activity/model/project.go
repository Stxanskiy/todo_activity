package model

import (
	"time"
)

type Project struct {
	ProjectID   int       `json:"project_id"`
	UserID      int       `json:"user_id"`
	ProjectName string    `json:"project_name"`
	CreateDate  time.Time `json:"create_date"`
}
