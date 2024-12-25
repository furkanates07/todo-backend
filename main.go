package main

import (
	"log"
	"net/http"
	"todo-backend/database"
)

func main() {
	database.ConnectDB()

	InitRoutes()

	log.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
