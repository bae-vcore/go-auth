package entity

import "time"

type User struct {
	ID         uint      `json:"id" gorm:"primaryKey,autoIncrement"`
	Email      string    `json:"email"`
	Password   string    `json:"-"`
	Name       string    `json:"name"`
	CreatedAt  time.Time `json:"created_at"`
	UpdateddAt time.Time `json:"updated_at"`
}
