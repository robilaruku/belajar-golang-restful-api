package helper

import (
	"robi/belajar-golang-restful-api/model/domain"
	"robi/belajar-golang-restful-api/model/web"
)

func ToCategoryResponse(category domain.Category) web.CategoryResponse {
	return web.CategoryResponse{
		Id: category.Id,
		Name: category.Name,
	}
}

func ToCategoryResponses(categories []domain.Category) []web.CategoryResponse  {
	var cateogiresResponses []web.CategoryResponse
	
	for _, cateogry := range categories {
		cateogiresResponses = append(cateogiresResponses, ToCategoryResponse(cateogry))
	}

	return cateogiresResponses
}