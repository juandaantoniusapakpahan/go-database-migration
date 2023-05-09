package main

import (
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/go-playground/validator/v10"
	"github.com/juandaantoniusapakpahan/go-restful-api/controller"
	"github.com/juandaantoniusapakpahan/go-restful-api/database"
	"github.com/juandaantoniusapakpahan/go-restful-api/exception"
	"github.com/juandaantoniusapakpahan/go-restful-api/helper"
	"github.com/juandaantoniusapakpahan/go-restful-api/repository"
	"github.com/juandaantoniusapakpahan/go-restful-api/services"
	"github.com/julienschmidt/httprouter"
)

func main() {
	db := database.Connect()
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := services.NewCategoryServie(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	router := httprouter.New()
	router.GET("/api/categories/:categoryId", categoryController.FindById)
	router.GET("/api/categories", categoryController.FindAll)
	router.POST("/api/categories", categoryController.Create)
	router.PUT("/api/categories/:categoryId", categoryController.Update)
	router.DELETE("/api/categories/:categoryId", categoryController.Delete)

	router.PanicHandler = exception.ErrorHalder
	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	fmt.Println("Server started at: localhost:8080")
	err := server.ListenAndServe()
	helper.ErrorHandle(err)
}
