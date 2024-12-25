package routes

import (
	"net/http"
)

func AuthRoutes() {
	http.HandleFunc("/register", RegisterHandler)
	http.HandleFunc("/login", LoginHandler)
}
