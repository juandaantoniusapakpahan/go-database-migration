package main

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/juandaantoniusapakpahan/go-restful-api/controller"
	"github.com/juandaantoniusapakpahan/go-restful-api/database"
	"github.com/juandaantoniusapakpahan/go-restful-api/repository"
	"github.com/juandaantoniusapakpahan/go-restful-api/services"
	"github.com/julienschmidt/httprouter"
)

func main() {
	db := database.Connect()
	validate := validator.New()
	newCategoryRepository := repository.NewCategoryRepository()
	newCategoryService := services.NewCategoryServie(newCategoryRepository, db, *validate)
	newCategoryController := controller.NewCategoryController(newCategoryService)

	router := httprouter.New()

	router.GET("/api/categories", newCategoryController.FindAll)
	router.GET("/api/categories/:categoryId", newCategoryController.FindById)
	router.DELETE("/api/categories/:categoryId", newCategoryController.Delete)
	router.POST("/api/categories", newCategoryController.Create)
	router.PUT("/api/categories/:categoryId", newCategoryController.Update)

	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	fmt.Println("Server started at: localhost:8080")
	server.ListenAndServe()
}
