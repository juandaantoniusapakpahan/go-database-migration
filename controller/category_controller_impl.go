package controller

import (
	"net/http"
	"strconv"

	"github.com/juandaantoniusapakpahan/go-restful-api/helper"
	"github.com/juandaantoniusapakpahan/go-restful-api/model/web"
	"github.com/juandaantoniusapakpahan/go-restful-api/services"
	"github.com/julienschmidt/httprouter"
)

type CategoryControllerImpl struct {
	CategoryService services.CategoryService
}

func NewCategoryController(categoryService services.CategoryService) CategoryController {
	return &CategoryControllerImpl{
		CategoryService: categoryService,
	}
}

func (controller *CategoryControllerImpl) Create(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	categoryCreateRequest := web.CategoryCreateRequest{}
	helper.RequestDecode(r, &categoryCreateRequest)

	result := controller.CategoryService.Create(r.Context(), categoryCreateRequest)

	response := web.Response{
		Code:   201,
		Status: "OK",
		Data:   result,
	}
	helper.ResponseJson(w, response)
}

func (controller *CategoryControllerImpl) Update(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	categoryUpdateRequest := web.CategoryUpdateRequest{}
	helper.RequestDecode(r, &categoryUpdateRequest)

	param := params.ByName("categoryId")
	categoryId, err := strconv.Atoi(param)
	helper.ErrorHandle(err)

	categoryUpdateRequest.Id = categoryId

	result := controller.CategoryService.Update(r.Context(), categoryUpdateRequest)

	response := web.Response{
		Code:   200,
		Status: "OK",
		Data:   result,
	}

	helper.ResponseJson(w, response)
}

func (controller *CategoryControllerImpl) Delete(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	param := params.ByName("categoryId")
	categoryId, err := strconv.Atoi(param)
	helper.ErrorHandle(err)

	controller.CategoryService.Delete(r.Context(), categoryId)

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

	result := controller.CategoryService.FindById(r.Context(), categoryId)

	response := web.Response{
		Code:   200,
		Status: "OK",
		Data:   result,
	}

	helper.ResponseJson(w, response)
}

func (controller *CategoryControllerImpl) FindAll(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	resutl := controller.CategoryService.FindAll(r.Context())

	webResponse := web.Response{
		Code:   200,
		Status: "OK",
		Data:   resutl,
	}

	helper.ResponseJson(w, webResponse)
}
