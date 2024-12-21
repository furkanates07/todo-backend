package database

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/supabase-community/supabase-go"
)

var SupabaseClient *supabase.Client

func ConnectDB() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	supabaseURL := os.Getenv("SUPABASE_URL")
	supabaseAPIKey := os.Getenv("SUPABASE_API_KEY")

	if supabaseURL == "" || supabaseAPIKey == "" {
		log.Fatal("Supabase URL or API key is missing from environment variables")
	}

	client, err := supabase.NewClient(supabaseURL, supabaseAPIKey, &supabase.ClientOptions{})
	if err != nil {
		log.Fatal("Error initializing Supabase client:", err)
	}

	SupabaseClient = client
	fmt.Println("Successfully connected to Supabase!")
}
