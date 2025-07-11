package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Employee struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	Role          string `json:"role"`
	Age           string `json:"age"`
	Qualification string `json:"qualification"`
}

type LoginInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func main() {
	router := gin.Default()

	// CORS configuration
	router.Use(func(c *gin.Context) {
		// Set CORS headers
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		// Handle OPTIONS request
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		// Set Origin header
		if origin := c.Request.Header.Get("Origin"); origin != "" {
			c.Writer.Header().Set("Origin", origin)
		}

		c.Next()
	})

	// Routes
	router.GET("/employees", getEmployees)
	router.POST("/login", loginUser)

	// Run server
	router.Run(":8082")
}

// GET /employees
func getEmployees(c *gin.Context) {
	employees := []Employee{
		{ID: 1, Name: "Suhail", Role: "Engineer", Age: "28", Qualification: "Bsc CS"},
		{ID: 2, Name: "Swalih", Role: "Manager", Age: "35", Qualification: "MBA"},
		{ID: 3, Name: "Ahmed", Role: "Designer", Age: "25", Qualification: "BFA"},
	}
	c.JSON(http.StatusOK, gin.H{"id": "1", "employees": employees})
}

// POST /login
func loginUser(c *gin.Context) {
	var input LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Dummy check
	if input.Username == "admin" && input.Password == "123" {
		c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid credentials"})
	}
}
