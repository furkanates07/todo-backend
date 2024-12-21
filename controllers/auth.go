package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"time"
	"todo-backend/database"
	"todo-backend/models"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var jwtKey = []byte(os.Getenv("JWT_SECRET"))

func Register(email, password string) (string, error) {
	client := database.SupabaseClient

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("error hashing password: %v", err)
	}

	user := models.User{
		Email:     email,
		Password:  string(hashedPassword),
		CreatedAt: models.MyTime(time.Now()),
	}

	_, _, err = client.From("users").Insert(user, false, "", "", "").Execute()
	if err != nil {
		return "", fmt.Errorf("error registering user: %v", err)
	}

	return "User registered successfully", nil
}

func Login(email, password string) (string, error) {
	client := database.SupabaseClient

	resp, _, err := client.From("users").Select("*", "exact", false).Eq("email", email).Single().Execute()
	if err != nil {
		return "", errors.New("user not found")
	}

	var user models.User
	if err := json.Unmarshal(resp, &user); err != nil {
		log.Println("Error unmarshalling user:", err)
		return "", errors.New("invalid user data")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", errors.New("invalid credentials")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", fmt.Errorf("error generating token: %v", err)
	}

	return tokenString, nil
}
