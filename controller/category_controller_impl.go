package controller

import (
	"context"
	"net/http"
	"strconv"

	"github.com/juandaantoniusapakpahan/go-restful-api/helper"
	"github.com/juandaantoniusapakpahan/go-restful-api/model/web"
	"github.com/juandaantoniusapakpahan/go-restful-api/services"
	"github.com/julienschmidt/httprouter"
)

type CategoryControllerImpl struct {
	Service services.CategoryService
}

func NewCategoryController(categoryService services.CategoryService) CategoryController {
	return &CategoryControllerImpl{
		Service: categoryService,
	}
}

func (controller *CategoryControllerImpl) Create(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	categoryCreateRequest := web.CategoryCreateRequest{}
	helper.RequestDecode(r, &categoryCreateRequest)

	result := controller.Service.Create(context.Background(), categoryCreateRequest)

	response := web.Response{
		Code:   200,
		Status: "OK",
		Data:   result,
	}
	helper.ResponseJson(w, response)
}

func (controller *CategoryControllerImpl) Update(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	categoryUpdateRequest := web.CategoryUpdateRequest{}
	helper.RequestDecode(r, &categoryUpdateRequest)

	result := controller.Service.Update(context.Background(), categoryUpdateRequest)

	responseData := web.Response{
		Code:   200,
		Status: "OK",
		Data:   result,
	}

	helper.ResponseJson(w, responseData)
}

func (controller *CategoryControllerImpl) Delete(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	param := params.ByName("categoryId")
	categoryId, err := strconv.Atoi(param)
	helper.ErrorHandle(err)

	controller.Service.Delete(context.Background(), categoryId)

	responesWeb := web.Response{
		Code:   200,
		Status: "OK",
	}

	helper.ResponseJson(w, responesWeb)

}

func (controller *CategoryControllerImpl) FindById(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	param := params.ByName("categoryId")
	categoryId, err := strconv.Atoi(param)
	helper.ErrorHandle(err)

	category := controller.Service.FindById(context.Background(), categoryId)

	webReponse := web.Response{
		Code:   200,
		Status: "OK",
		Data:   category,
	}

	helper.ResponseJson(w, webReponse)
}

func (controller *CategoryControllerImpl) FindAll(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	resutl := controller.Service.FindAll(context.Background())

	webResponse := web.Response{
		Code:   200,
		Status: "OK",
		Data:   resutl,
	}

	helper.ResponseJson(w, webResponse)
}
