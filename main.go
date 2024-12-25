package main

import (
	"log"
	"net/http"
	"todo-backend/database"
	authroutes "todo-backend/routes/auth"
	todoroutes "todo-backend/routes/todo"
)

func main() {
	database.ConnectDB()

	authroutes.AuthRoutes()
	todoroutes.TodoRoutes()

	log.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
