package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type certificate struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Name     string `json:"name"`
	Position int    `json:"position"`
}

var certificates = []certificate{
	{ID: "1", Title: "Web Development", Name: "Prathviraj", Position: 1},
	{ID: "2", Title: "App Development", Name: "Virat", Position: 2},
	{ID: "3", Title: "UI/UX Design", Name: "Rohit", Position: 1},
}

func main() {
	router := gin.Default()

	router.GET("/certificates", getCertificates)
	router.GET("/certificates/:id", getCertificateByID)
	router.POST("/certificates", postCertificate)
	router.PUT("/certificates/:id", updateCertificate) // Added Update Route

	router.Run("localhost:8080")
}

// Get all certificates
func getCertificates(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, certificates)
}

// Get certificate by ID
func getCertificateByID(c *gin.Context) {
	id := c.Param("id")
	for _, cert := range certificates {
		if cert.ID == id {
			c.IndentedJSON(http.StatusOK, cert)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Certificate not found"})
}

// Create a new certificate
func postCertificate(c *gin.Context) {
	var newCertificate certificate
	if err := c.BindJSON(&newCertificate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	certificates = append(certificates, newCertificate)
	c.IndentedJSON(http.StatusCreated, newCertificate)
}

// Update an existing certificate by ID
func updateCertificate(c *gin.Context) {
	id := c.Param("id")
	var updatedCertificate certificate

	if err := c.BindJSON(&updatedCertificate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	for i, cert := range certificates {
		if cert.ID == id {
			certificates[i] = updatedCertificate
			c.IndentedJSON(http.StatusOK, updatedCertificate)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Certificate not found"})
}

