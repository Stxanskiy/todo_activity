package model

import "time"

type Task struct {
	TaskID      int       `json:"task_id"`
	ProjectID   int       `json:"project_id"`
	TaskName    string    `json:"task_name"`
	Description string    `json:"description"`
	DueDate     time.Time `json:"due_date"`
	Priority    string    `json:"priority"`
	Status      string    `json:"status"`
	CreateDate  time.Time `json:"create_date"`
}
