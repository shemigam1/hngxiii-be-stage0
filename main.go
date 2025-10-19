package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

// User represents the user information
type User struct {
	Email string `json:"email"`
	Name  string `json:"name"`
	Stack string `json:"stack"`
}

// CatFactResponse represents the response from the Cat Facts API
type CatFactResponse struct {
	Fact string `json:"fact"`
}

// MeResponse represents the response structure for /me endpoint
type MeResponse struct {
	Status    string    `json:"status"`
	User      User      `json:"user"`
	Timestamp string    `json:"timestamp"`
	Fact      string    `json:"fact"`
}

// getCatFact fetches a random cat fact from the Cat Facts API
func getCatFact() string {
	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	resp, err := client.Get("https://catfact.ninja/fact")
	if err != nil {
		log.Printf("Error fetching cat fact: %v", err)
		return "Cats have over 20 different vocal sounds they use to communicate with each other and humans."
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("Cat Facts API returned status: %d", resp.StatusCode)
		return "Cats have over 20 different vocal sounds they use to communicate with each other and humans."
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error reading response body: %v", err)
		return "Cats have over 20 different vocal sounds they use to communicate with each other and humans."
	}

	var catFact CatFactResponse
	err = json.Unmarshal(body, &catFact)
	if err != nil {
		log.Printf("Error unmarshaling cat fact: %v", err)
		return "Cats have over 20 different vocal sounds they use to communicate with each other and humans."
	}

	return catFact.Fact
}

func main() {
	StartGin()
}

// StartGin starts gin web server with setting router.
func StartGin() {
	gin.SetMode(gin.ReleaseMode)

	router := gin.New()
	router.Use(gin.Recovery())

	// Health check endpoint
	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	// Profile endpoint
	router.GET("/me", func(c *gin.Context) {
		c.Header("Content-Type", "application/json")

		email := "semiloreomotade@gmail.com"
		name := "Oluwasemilore Omotade-Michaels"
		stack := "Go/Gin"

		fact := getCatFact()

		response := MeResponse{
			Status: "success",
			User: User{
				Email: email,
				Name:  name,
				Stack: stack,
			},
			Timestamp: time.Now().UTC().Format(time.RFC3339Nano),
			Fact:      fact,
		}

		c.JSON(http.StatusOK, response)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Starting server on port %s", port)
	if err := router.Run(":" + port); err != nil {
		log.Panicf("error: %s", err)
	}
}