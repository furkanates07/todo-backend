package main

import (
	"log"
	"net/http"
	"todo-backend/database"
	"todo-backend/routes"
)

func main() {
	database.ConnectDB()

	routes.AuthRoutes()
	routes.TodoRoutes()

	log.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
