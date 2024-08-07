package model

import "time"

type Comment struct {
	CommentID   int       `json:"comment_id"`
	TaskID      int       `json:"task_id"`
	UserID      int       `json:"user_id"`
	CommetnText string    `json:"commetn_text"`
	CreateDate  time.Time `json:"create_date"`
}
