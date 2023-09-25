package helper

import (
	"mymodule/model/domain"
	"mymodule/model/web"
)

func ToCategoryResponse(category domain.Category) web.CategoryResponse {

	return web.CategoryResponse{
		Id: category.Id,
		Name: category.Name,
	}
}


func ToCategoriesResponse(categories []domain.Category) []web.CategoryResponse {

	categoriesResponse := []web.CategoryResponse{}
	for _, category := range categories {
		
		categoriesResponse = append(categoriesResponse, ToCategoryResponse(category))
		
	}

	return categoriesResponse
}