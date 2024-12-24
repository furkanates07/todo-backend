package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"
	"todo-backend/database"
	"todo-backend/models"

	"github.com/google/uuid"
)

func CreateTodo(user models.User, title, description string) (models.Todo, error) {
	client := database.SupabaseClient

	todo := models.Todo{
		ID:          uuid.New(),
		UserID:      user.ID,
		Title:       title,
		Description: description,
		Status:      models.PENDING,
		CreatedAt:   models.MyTime(time.Now()),
	}

	fmt.Println("Inserting new todo for user:", user.ID)

	resp, _, err := client.From("todos").Insert(todo, false, "", "", "").Execute()
	if err != nil || len(resp) == 0 {
		return models.Todo{}, fmt.Errorf("error creating todo: %v", err)
	}

	fmt.Println("Todo created successfully!")
	return todo, nil
}

func GetTodos(user models.User) ([]models.Todo, error) {
	client := database.SupabaseClient

	resp, _, err := client.From("todos").Select("*", "exact", false).Eq("user_id", user.ID.String()).Execute()
	if err != nil {
		return nil, fmt.Errorf("error fetching todos: %v", err)
	}

	var todos []models.Todo
	if err := json.Unmarshal(resp, &todos); err != nil {
		return nil, fmt.Errorf("error unmarshalling todos: %v", err)
	}

	return todos, nil
}

func UpdateTodo(user models.User, todoID uuid.UUID, title, description string, status models.TodoStatus) (models.Todo, error) {
	client := database.SupabaseClient

	updateData := map[string]interface{}{
		"title":       title,
		"description": description,
		"status":      status,
	}

	_, _, err := client.From("todos").Update(updateData, "exact", "").Eq("id", todoID.String()).Eq("user_id", user.ID.String()).Execute()
	if err != nil {
	}

	return models.Todo{
		ID:          todoID,
		UserID:      user.ID,
		Title:       title,
		Description: description,
		Status:      status,
		CreatedAt:   models.MyTime(time.Now()),
	}, nil
}

func DeleteTodo(user models.User, todoID uuid.UUID) error {
	client := database.SupabaseClient

	_, _, err := client.From("todos").Delete("exact", "").Eq("id", todoID.String()).Eq("user_id", user.ID.String()).Execute()
	if err != nil {
		return fmt.Errorf("error deleting todo: %w", err)
	}

	return nil
}
