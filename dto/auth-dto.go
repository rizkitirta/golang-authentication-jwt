package dto

type LoginDTO struct {
	Email string `json:"email" binding:"required" form:"email" validate:"email"`
	Password string `json:"password" binding:"required" form:"password" validate:"min=8"`
}
 
