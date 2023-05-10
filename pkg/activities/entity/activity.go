package entityActivity

import "time"

type Activity struct {
	ActivityID int64     `json:"id" gorm:"primaryKey"`
	Title      string    `json:"title"`
	Email      string    `json:"email"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func NewActivity(title string, email string) Activity {
	return Activity{
		Title:     title,
		Email:     email,
		CreatedAt: time.Now(),
	}
}
