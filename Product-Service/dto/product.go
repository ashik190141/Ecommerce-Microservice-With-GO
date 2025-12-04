package dto

type CreateProductDTO struct {
	Name  string  `json:"name" validate:"required"`
	Price float64 `json:"price" validate:"required,gt=0"`
	Stock int     `json:"stock" validate:"required,gte=0"`
	Email string  `json:"userEmail" db:"useremail"`
	Sku   string  `json:"sku" validate:"required"`
}

type UpdateProductDTO struct {
	Name  string  `json:"name,omitempty"`
	Price float64 `json:"price,omitempty" validate:"omitempty,gt=0"`
	Stock int     `json:"stock,omitempty" validate:"omitempty,gte=0"`
	Sku   string  `json:"sku,omitempty"`
}

type GetProductResponse struct {
	Id int `json:"id" validate:"required"`
	CreateProductDTO
	CreatedAt string `json:"createdAt" db:"created_at"`
	UpdatedAt string `json:"updatedAt" db:"updated_at"`
}
