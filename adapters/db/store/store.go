package store

import (
	"github.com/Liberkeys/api-skeleton/ports/models"
)

// UserStore explicit store
type UserStore interface {
	SelectAll() ([]models.User, error) // Gateway approach
}
