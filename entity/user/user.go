package user

import (
	"time"
)

type NewUserInput struct {
	Email        string `json:"email" binding:"required"`
	FullName     string `json:"fullname" binding:"required"`
	Password     string `json:"password" binding:"required"`
	PhoneNumber  string `json:"phone_number" binding:"required"`
	Sex          int    `json:"sex" binding:"required"`
	Biography    string `json:"biography" binding:"required"`
	Location     string `json:"location" binding:"required"`
	DateOfBirth  string `json:"date_of_birth" binding:"required"`
	ProfilePhoto string `json:"profile_photo" binding:"required"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	Token             string `json:"token"`
	IsPasswordCorrect bool   `json:"is_password_correct"`
}

type User struct {
	ID           int64     `gorm:"column:id;primaryKey"`
	Email        string    `json:"email" gorm:"column:email"`
	FullName     string    `json:"fullname" gorm:"column:fullname"`
	Password     string    `json:"password" gorm:"column:password"`
	PhoneNumber  string    `json:"phone_number" gorm:"column:phone_number"`
	Sex          int       `json:"sex" gorm:"column:sex"`
	Biography    string    `json:"biography" gorm:"column:biography"`
	Location     string    `json:"location" gorm:"column:location"`
	DateOfBirth  time.Time `json:"date_of_birth" gorm:"column:date_of_birth"`
	ProfilePhoto string    `json:"profile_photo" gorm:"column:profile_photo"`
	CreatedAt    time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt    time.Time `json:"updated_at" gorm:"column:updated_at"`
}

func (User) TableName() string {
	return "useraccount"
}
