package service

import (
	"errors"
	"pizzaria/internal/models"
)

func ValidatePizza(pizza *models.Pizza) error {
	if pizza.Preco < 0 {
		return errors.New("o preço da pizza não pode ser negativo")
	}
	return nil
}
