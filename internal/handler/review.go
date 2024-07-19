package handler

import (
	"net/http"
	"pizzaria/internal/data"
	"pizzaria/internal/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func PostReview(c *gin.Context) {
	pizzaID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	var newReview models.Review
	if err := c.ShouldBindJSON(&newReview); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	for i, p := range data.Pizzas {
		if p.ID == pizzaID {
			p.Review = append(p.Review, newReview)
			data.Pizzas[i] = p
			data.SavePizza()
			c.JSON(http.StatusCreated, p)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "Pizza not found"})
}
