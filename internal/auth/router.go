package auth

import (
	"net/http"
)

func InitRoutes() {
	http.HandleFunc("/register", RegisterHandler)
	http.HandleFunc("/login", LoginHandler)
}
