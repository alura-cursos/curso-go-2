package service

import (
	"errors"
	"pizzaria/internal/models"
)

func ValidateReviewRating(review models.Review) error {
	if review.Rating < 1 || review.Rating > 5 {
		return errors.New("rating must be between 1 and 5")
	}
	return nil
}
