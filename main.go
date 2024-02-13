package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Car struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Color string `json:"color"`
}

var Cars []Car

func main() {
	Cars = []Car{
		{ID: "1", Title: "BMW", Color: "Black"},
		{ID: "2", Title: "Tesla", Color: "Red"},
	}

	router := gin.Default()

	router.GET("/cars", getCars)
	router.GET("/cars/:id", getCarByID)
	router.POST("/cars", createCar)
	router.DELETE("/cars/:id", deleteCar)

	err := router.Run("localhost:8080")
	if err != nil {
		return
	}
}

func createCar(c *gin.Context) {
	var newCar Car
	if err := c.BindJSON(&newCar); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Failed to create a car"})
		return
	}

	Cars = append(Cars, newCar)
	c.IndentedJSON(http.StatusCreated, newCar)
}

func getCars(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, Cars)
}

func getCarByID(c *gin.Context) {
	for _, car := range Cars {
		if car.ID == c.Param("id") {
			c.IndentedJSON(http.StatusOK, car)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Requested car is not found"})
}

func deleteCar(c *gin.Context) {
	for index, car := range Cars {
		if car.ID == c.Param("id") {
			Cars = append(Cars[:index], Cars[index+1:]...)
			c.IndentedJSON(http.StatusOK, gin.H{"message": "Car is deleted"})
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Car is not found"})
}
