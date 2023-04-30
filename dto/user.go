package dto

import "time"

type RegisterRequest struct {
	Fullname string `json:"full_name" valid:"required~full_name cannot be empty"`
	Email    string `json:"email" valid:"required~email cannot be empty"`
	Password string `json:"password" valid:"required~password cannot be empty,minstringlength(6)~password must be at least 6 characters long"`
}

type RegisterResponse struct {
	ID        uint      `json:"id"`
	Fullname  string    `json:"full_name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}
