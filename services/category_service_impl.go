package services

import (
	"context"
	"database/sql"

	"github.com/juandaantoniusapakpahan/go-restful-api/helper"
	"github.com/juandaantoniusapakpahan/go-restful-api/model/domain"
	"github.com/juandaantoniusapakpahan/go-restful-api/model/web"
	"github.com/juandaantoniusapakpahan/go-restful-api/repository"
)

type CategoryServiceImpl struct {
	Repository repository.CategoryRepository
	DB         *sql.DB
}

func (service *CategoryServiceImpl) Save(ctx context.Context, category web.CategoryCreateRequest) web.CategoryResponse {
	tx, err := service.DB.Begin()
	helper.ErrorHandle(err)
	helper.CommitRollBack(tx)

	domainCategory := domain.Category{
		Name: category.Name,
	}
	result := service.Repository.Save(ctx, tx, domainCategory)

	return helper.CategoryToResponse(result)
}
