package main

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"

	"robi/belajar-golang-restful-api/app"
	"robi/belajar-golang-restful-api/controller"
	"robi/belajar-golang-restful-api/helper"
	"robi/belajar-golang-restful-api/middleware"
	"robi/belajar-golang-restful-api/repository"
	"robi/belajar-golang-restful-api/service"
)

func main() {

	db := app.NewDB()
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	router := app.NewRouter(categoryController)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
