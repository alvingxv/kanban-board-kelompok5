package dto

import "time"

type CategoryRequest struct {
	Type string `json:"type" valid:"required~type cannot be empty"`
}

type CreateCategoryResponse struct {
	Id        uint      `json:"id"`
	Type      string    `json:"type"`
	CreatedAt time.Time `json:"created_at"`
}
type UpdateCategoryResponse struct {
	Id        uint      `json:"id"`
	Type      string    `json:"type"`
	UpdatedAt time.Time `json:"updated_at"`
}
