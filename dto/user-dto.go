package dto

type UserCreateDTO struct {
	Name        string `json:"name" binding:"required" form:"name"`
	Email       string `json:"email" binding:"required" form:"email" validate:"email"`
	Password    string `json:"password" binding:"required" form:"password" validate:"min=8"`
}

type UserUpdateDTO struct {
	ID          int    `json:"id" binding:"required" form:"id"`
	Name        string `json:"name" binding:"required" form:"name"`
	Email       string `json:"email" binding:"required" form:"email" validate:"email"`
	Password    string `json:"password,omitempty" form:"password,omitempty" validate:"min=8"`
}