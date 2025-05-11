package dto

type CreateProductInput struct {
	Name  string  `json:"name" validate:"required"`
	Price float64 `json:"price" validate:"required"`
}

type CreateUserInput struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}
