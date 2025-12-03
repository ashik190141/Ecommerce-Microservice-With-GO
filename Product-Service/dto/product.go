package dto

type CreateProductDTO struct {
	Name  string  `json:"name" validate:"required"`
	Price float64 `json:"price" validate:"required,gt=0"`
	Stock int     `json:"stock" validate:"required,gte=0"`
	Email string  `json:"userEmail" db:"useremail"`
	Sku   string  `json:"sku" validate:"required"`
}

type GetProductResponse struct {
	Id int `json:"id" validate:"required"`
	CreateProductDTO
	CreatedAt string `json:"createdAt" db:"created_at"`
	UpdatedAt string `json:"updatedAt" db:"updated_at"`
}
