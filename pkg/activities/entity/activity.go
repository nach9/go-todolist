package entityActivity

import "time"

type Activity struct {
	ActivityID int64     `json:"id" gorm:"primaryKey"`
	Title      string    `json:"title"`
	Email      string    `json:"email"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}

func NewActivity(title string, email string) Activity {
	return Activity{
		Title: title,
		Email: email,
	}
}
