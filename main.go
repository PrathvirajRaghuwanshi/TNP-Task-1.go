package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type certificate struct {
	ID       string  `json:"id"`
	Title    string  `json:"title"`
	Name     string  `json:"name"`
	Position float64 `json:"position"`
}

var certificates = []certificate{
	{ID: "1", Title: "web devlopment", Name: "Prathviraj", Position: 1},
	{ID: "2", Title: "app devlopment", Name: "Virat", Position: 2},
	{ID: "3", Title: "UI/UX designer", Name: "Rohit", Position: 1},
}

func main() {
	router := gin.Default()
	router.GET("/certificates", getcertificate)
	router.GET("/certificates/:id", getcertificateByID)
	router.POST("/certificates", postcertificate)

	router.Run("localhost:8080")
}

func getcertificate(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, certificates)
}

func postcertificate(c *gin.Context) {
	var newCertificate certificate

	if err := c.BindJSON(&newCertificate); err != nil {
		return
	}

	certificates = append(certificates, newCertificate)
	c.IndentedJSON(http.StatusCreated, newCertificate)
}

func getcertificateByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range certificates {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "certificate not found"})
}
