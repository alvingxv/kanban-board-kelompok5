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

type LoginRequest struct {
	Email    string `json:"email" valid:"required~email cannot be empty"`
	Password string `json:"password" valid:"required~password cannot be empty"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type UpdateRequest struct {
	Fullname string `json:"full_name"`
	Email    string `json:"email"`
}

type UpdateResponse struct {
	ID        uint      `json:"id"`
	Fullname  string    `json:"full_name"`
	Email     string    `json:"email"`
	UpdatedAt time.Time `json:"updated_at"`
}
