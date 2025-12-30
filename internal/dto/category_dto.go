package dtos

type CategoryRequest struct {
	Name string `json:"name" binding:"required"`
}

type CategoryResponse struct {
	ID  string `json:"id"`
	Name string `json:"name"`
}

type ListCategoriesResponse struct {
	Categories []CategoryResponse `json:"categories"`
}
