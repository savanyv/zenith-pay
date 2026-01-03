package dtos

type ProductRequest struct {
	CategoryID string  `json:"category_id" validate:"required,uuid"`
	SKU     string  `json:"sku"`
	Name	 string  `json:"name" validate:"required"`
	Price   float64 `json:"price" validate:"required,gt=0"`
	Stock   int     `json:"stock" validate:"required,gte=0"`
}

type ProductResponse struct {
	ID         string  `json:"id"`
	CategoryID string  `json:"category_id"`
	CategoryName string `json:"category_name"`
	SKU        string  `json:"sku"`
	Name       string  `json:"name"`
	Price      float64 `json:"price"`
	Stock      int     `json:"stock"`
	CreatedAt  string  `json:"created_at"`
	UpdatedAt  string  `json:"updated_at"`
}

type ProductUpdateRequest struct {
	CategoryID *string `json:"category_id,omitempty" validate:"omitempty,uuid"`
	Name      *string `json:"name,omitempty"`
	Price     *float64 `json:"price,omitempty" validate:"omitempty,gt=0"`
	Stock     *int    `json:"stock,omitempty" validate:"omitempty,gte=0"`
}
