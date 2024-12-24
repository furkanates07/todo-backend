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
	if err != nil {
		return models.Todo{}, fmt.Errorf("error creating todo: %w", err)
	}

	if len(resp) == 0 {
		return models.Todo{}, errors.New("no response from database")
	}

	fmt.Println("Todo created successfully!")
	return todo, nil
}

func GetTodos(user models.User) ([]models.Todo, error) {
	client := database.SupabaseClient

	resp, _, err := client.From("todos").Select("*", "exact", false).Eq("user_id", user.ID.String()).Execute()
	if err != nil {
		return nil, fmt.Errorf("error fetching todos: %w", err)
	}

	if len(resp) == 0 {
		return nil, errors.New("no todos found")
	}

	var todos []models.Todo
	if err := json.Unmarshal(resp, &todos); err != nil {
		return nil, fmt.Errorf("error unmarshalling todos: %w", err)
	}

	return todos, nil
}

func GetTodo(todoID uuid.UUID) (models.Todo, error) {
	client := database.SupabaseClient

	resp, _, err := client.From("todos").Select("*", "exact", false).Eq("id", todoID.String()).Single().Execute()
	if err != nil {
		return models.Todo{}, fmt.Errorf("error fetching todo: %w", err)
	}

	if len(resp) == 0 {
		return models.Todo{}, errors.New("todo not found")
	}

	var fetchedTodo models.Todo
	if err := json.Unmarshal(resp, &fetchedTodo); err != nil {
		return models.Todo{}, fmt.Errorf("error unmarshalling todo: %w", err)
	}

	return fetchedTodo, nil
}

func UpdateTodo(todo models.Todo) (models.Todo, error) {
	client := database.SupabaseClient

	resp, _, err := client.From("todos").Update(todo, "exact", "").Eq("id", todo.ID.String()).Execute()
	if err != nil {
		return models.Todo{}, fmt.Errorf("error updating todo: %w", err)
	}

	if len(resp) == 0 {
		return todo, nil
	}

	var updatedTodo models.Todo
	if err := json.Unmarshal(resp, &updatedTodo); err != nil {
		return models.Todo{}, fmt.Errorf("error unmarshalling updated todo: %w", err)
	}

	return updatedTodo, nil
}

func UpdateTodoStatus(todoID uuid.UUID, status models.TodoStatus) (models.Todo, error) {
	client := database.SupabaseClient

	updateData := map[string]interface{}{
		"status": status,
	}

	resp, _, err := client.From("todos").Update(updateData, "exact", "").Eq("id", todoID.String()).Execute()
	if err != nil {
		return models.Todo{}, fmt.Errorf("error updating todo status: %w", err)
	}

	if len(resp) == 0 {
		todo, err := GetTodo(todoID)
		if err != nil {
			return models.Todo{}, err
		}
		return todo, nil
	}

	var updatedTodo models.Todo
	if err := json.Unmarshal(resp, &updatedTodo); err != nil {
		return models.Todo{}, fmt.Errorf("error unmarshalling updated todo: %w", err)
	}

	return updatedTodo, nil
}

func DeleteTodo(todoID uuid.UUID) error {
	client := database.SupabaseClient

	_, _, err := client.From("todos").Delete("exact", "").Eq("id", todoID.String()).Execute()
	if err != nil {
		return fmt.Errorf("error deleting todo: %w", err)
	}

	fmt.Println("Todo deleted successfully!")

	return nil
}
