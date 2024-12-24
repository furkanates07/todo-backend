package routes

import (
	"encoding/json"
	"net/http"
	"todo-backend/controllers"
	"todo-backend/middlewares"
	"todo-backend/models"

	"github.com/google/uuid"
)

func CreateTodoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	userID, ok := r.Context().Value(middlewares.UserIDKey).(string)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	var requestBody struct {
		Title       string `json:"title"`
		Description string `json:"description"`
	}

	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if userID != requestBody.UserID {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	user := models.User{ID: uuid.MustParse(userID)}
func GetTodoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	userID, ok := r.Context().Value(middlewares.UserIDKey).(uuid.UUID)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	todoIDStr := r.URL.Path[len("/todo/get/"):]
	todoID, err := uuid.Parse(todoIDStr)
	if err != nil {
		http.Error(w, "Invalid Todo ID", http.StatusBadRequest)
		return
	}

	todo, err := controllers.GetTodo(todoID)
	if err != nil {
		if err.Error() == "todo not found" {
			http.Error(w, "Todo not found", http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if todo.UserID != userID {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todo)
}

func GetTodosHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	userID, ok := r.Context().Value(middlewares.UserIDKey).(string)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	user := models.User{ID: uuid.MustParse(userID)}

	todos, err := controllers.GetTodos(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}

func TodoRoutes() {
	http.Handle("/todo/create", middlewares.AuthMiddleware(http.HandlerFunc(CreateTodoHandler)))

	http.Handle("/todo/get/", middlewares.AuthMiddleware(http.HandlerFunc(GetTodoHandler)))

	http.Handle("/todo/get", middlewares.AuthMiddleware(http.HandlerFunc(GetTodosHandler)))
}
