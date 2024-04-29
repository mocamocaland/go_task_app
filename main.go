package main

import (
	"go_task_app/controller"
	"go_task_app/db"
	"go_task_app/repository"
	"go_task_app/router"
	"go_task_app/usecase"
)

func main() {
	db := db.NewDB()
	userRepository := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository)
	userController := controller.NewUserController(userUsecase)
	e := router.NewRouter(userController)
	e.Logger.Fatal(e.Start(":8080"))
}
