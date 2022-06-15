package dto

type BookUpdateDTO struct {
	ID          int    `json:"id" binding:"required" form:"id"`
	Title       string `json:"title" binding:"required" form:"title"`
	Description string `json:"description" binding:"required" form:"description"`
	UserID      int    `json:"user_id,omitempty" form:"user_id,omitempty"`
}

type BookCreateDTO struct {
	Title       string `json:"title" binding:"required" form:"title"`
	Description string `json:"description" binding:"required" form:"description"`
	UserID      int    `json:"user_id,omitempty" form:"user_id,omitempty"`
}
