package model

import "time"

type Reminder struct {
	ReminderID   int       `json:"reminder_id"`
	TaskID       int       `json:"task_id"`
	ReminderDate time.Time `json:"reminder_date"`
}
