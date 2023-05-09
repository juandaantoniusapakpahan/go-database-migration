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
	CategoryRepository repository.CategoryRepository
	DB                 *sql.DB
	Validate           *validator.Validate
}

func NewCategoryServie(respotiry repository.CategoryRepository, db *sql.DB, validate *validator.Validate) CategoryService {
	return &CategoryServiceImpl{
		CategoryRepository: respotiry,
		DB:                 db,
		Validate:           validate,
	}
}

func (service *CategoryServiceImpl) Create(ctx context.Context, category web.CategoryCreateRequest) web.CategoryResponse {
	err := service.Validate.Struct(category)
	helper.ErrorHandle(err)

	tx, err := service.DB.Begin()
	helper.ErrorHandle(err)
	defer helper.CommitRollBack(tx)

	domainCategory := domain.Category{
		Name: category.Name,
	}
	result := service.CategoryRepository.Save(ctx, tx, domainCategory)

	return helper.CategoryToResponse(result)
}

func (service *CategoryServiceImpl) Update(ctx context.Context, request web.CategoryUpdateRequest) web.CategoryResponse {
	err := service.Validate.Struct(request)
	helper.ErrorHandle(err)

	tx, err := service.DB.Begin()
	helper.ErrorHandle(err)
	defer helper.CommitRollBack(tx)

	domainCategory, err := service.CategoryRepository.FindById(ctx, tx, request.Id)

	domainCategory.Name = request.Name

	category := service.CategoryRepository.Update(ctx, tx, domainCategory)

	return helper.CategoryToResponse(category)
}

func (service *CategoryServiceImpl) Delete(ctx context.Context, categoryId int) {
	tx, err := service.DB.Begin()
	helper.ErrorHandle(err)
	defer helper.CommitRollBack(tx)

	domainCategory, err := service.CategoryRepository.FindById(ctx, tx, categoryId)
	helper.ErrorHandle(err)

	service.CategoryRepository.Delete(ctx, tx, domainCategory)
}

func (service *CategoryServiceImpl) FindById(ctx context.Context, categoryId int) web.CategoryResponse {
	tx, err := service.DB.Begin()
	helper.ErrorHandle(err)
	defer helper.CommitRollBack(tx)

	domainCategory, err := service.CategoryRepository.FindById(ctx, tx, categoryId)

	return helper.CategoryToResponse(domainCategory)
}

func (service *CategoryServiceImpl) FindAll(ctx context.Context) []web.CategoryResponse {
	tx, err := service.DB.Begin()
	helper.ErrorHandle(err)
	defer helper.CommitRollBack(tx)

	categories := service.CategoryRepository.FindAll(ctx, tx)

	return helper.CategoriesToResponses(categories)

}
