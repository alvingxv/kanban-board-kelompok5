package dto

import "time"

type CreateCategoryRequest struct {
	Type string `json:"type" valid:"required~type cannot be empty"`
}

type CreateCategoryResponse struct {
	Id        uint      `json:"id"`
	Type      string    `json:"type"`
	CreatedAt time.Time `json:"created_at"`
}
