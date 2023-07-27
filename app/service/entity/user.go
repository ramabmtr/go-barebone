package entity

import "time"

type User struct {
	Username  string    `json:"username" gorm:"primaryKey"`
	CreatedAt time.Time `json:"created_at"`
}

type UserRegisterParam struct {
	Username string `json:"username" validate:"required,min=6"`
}

type UserLoginParam struct {
	Username string `json:"username"`
}

type UserLoginResponse struct {
	Token string `json:"token"`
}
