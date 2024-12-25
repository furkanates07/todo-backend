package main

import (
	"todo-backend/internal/auth"
	"todo-backend/internal/todo"
)

func InitRoutes() {
	auth.InitRoutes()
	todo.InitRoutes()
}
