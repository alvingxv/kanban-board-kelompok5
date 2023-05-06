package dto

import "time"

type CreateTaskRequest struct {
	Title       string `json:"title" valid:"required~title cannot be empty"`
	Description string `json:"description" valid:"required~description cannot be empty"`
	CategoryId  uint   `json:"category_id" valid:"required~category_id cannot be empty"`
}

type CreateTaskResponse struct {
	ID          uint      `json:"id"`
	Title       string    `json:"title"`
	Status      bool      `json:"status"`
	Description string    `json:"description"`
	UserID      uint      `json:"user_id"`
	CategoryID  uint      `json:"category_id"`
	CreatedAt   time.Time `json:"created_at"`
}
