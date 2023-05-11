package entityTodo

import "time"

type Todo struct {
	TodoID          int64     `json:"id" gorm:"primaryKey"`
	ActivityGroupID int64     `json:"activity_group_id"`
	Title           string    `json:"title"`
	Priority        string    `json:"priority"`
	IsActive        bool      `json:"is_active"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updatedAt"`
}
