package services

import (
	"context"

	"github.com/juandaantoniusapakpahan/go-restful-api/model/web"
)

type CategoryService interface {
	Create(ctx context.Context, category web.CategoryCreateRequest) web.CategoryResponse
	Update(ctx context.Context, category web.CategoryUpdateRequest) web.CategoryResponse
	Delete(ctx context.Context, categoryId string)
	FindById(ctx context.Context, categoryId string) web.CategoryResponse
	FindAll(ctx context.Context) []web.CategoryResponse
}
