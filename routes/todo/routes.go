package routes

import (
	"net/http"
	"todo-backend/middlewares"
)

func TodoRoutes() {
	http.Handle("/todo/create", middlewares.AuthMiddleware(http.HandlerFunc(CreateTodoHandler)))

	http.Handle("/todo/get/", middlewares.AuthMiddleware(http.HandlerFunc(GetTodoHandler)))

	http.Handle("/todo/get", middlewares.AuthMiddleware(http.HandlerFunc(GetTodosHandler)))

	http.Handle("/todo/update/", middlewares.AuthMiddleware(http.HandlerFunc(UpdateTodoHandler)))

	http.Handle("/todo/update/status/", middlewares.AuthMiddleware(http.HandlerFunc(UpdateTodoStatusHandler)))

	http.Handle("/todo/delete/", middlewares.AuthMiddleware(http.HandlerFunc(DeleteTodoHandler)))
}
