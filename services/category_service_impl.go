package services

import (
	"context"
	"database/sql"

	"github.com/go-playground/validator/v10"
	"github.com/juandaantoniusapakpahan/go-restful-api/helper"
	"github.com/juandaantoniusapakpahan/go-restful-api/model/domain"
	"github.com/juandaantoniusapakpahan/go-restful-api/model/web"
	"github.com/juandaantoniusapakpahan/go-restful-api/repository"
)

type CategoryServiceImpl struct {
	Repository repository.CategoryRepository
	DB         *sql.DB
	Validate   *validator.Validate
}

func NewCategoryServie(respotiry repository.CategoryRepository, db *sql.DB, validate validator.Validate) CategoryService {
	return &CategoryServiceImpl{
		Repository: respotiry,
		DB:         db,
		Validate:   &validate,
	}
}

func (service *CategoryServiceImpl) Create(ctx context.Context, category web.CategoryCreateRequest) web.CategoryResponse {
	err := service.Validate.Struct(category)
	tx, err := service.DB.Begin()
	helper.ErrorHandle(err)
	defer helper.CommitRollBack(tx)

	domainCategory := domain.Category{
		Name: category.Name,
	}
	result := service.Repository.Save(ctx, tx, domainCategory)

	return helper.CategoryToResponse(result)
}

func (service *CategoryServiceImpl) Update(ctx context.Context, category web.CategoryUpdateRequest) web.CategoryResponse {
	err := service.Validate.Struct(category)

	tx, err := service.DB.Begin()
	helper.ErrorHandle(err)
	defer helper.CommitRollBack(tx)

	domainCategory, err := service.Repository.FindById(ctx, tx, category.Id)

	helper.ErrorHandle(err)

	domainCategory.Name = category.Name

	domainCategoryResult := service.Repository.Update(ctx, tx, domainCategory)

	return helper.CategoryToResponse(domainCategoryResult)
}

func (service *CategoryServiceImpl) Delete(ctx context.Context, categoryId int) {
	tx, err := service.DB.Begin()
	helper.ErrorHandle(err)
	defer helper.CommitRollBack(tx)

	domainCategory, err := service.Repository.FindById(ctx, tx, categoryId)
	helper.ErrorHandle(err)

	service.Repository.Delete(ctx, tx, domainCategory)
}

func (service *CategoryServiceImpl) FindById(ctx context.Context, categoryId int) web.CategoryResponse {
	tx, err := service.DB.Begin()
	helper.ErrorHandle(err)
	defer helper.CommitRollBack(tx)

	domainCategory, err := service.Repository.FindById(ctx, tx, categoryId)
	helper.ErrorHandle(err)

	return helper.CategoryToResponse(domainCategory)
}

func (service *CategoryServiceImpl) FindAll(ctx context.Context) []web.CategoryResponse {
	tx, err := service.DB.Begin()
	helper.ErrorHandle(err)
	defer helper.CommitRollBack(tx)

	categories := service.Repository.FindAll(ctx, tx)

	return helper.CategoriesToResponses(categories)

}
