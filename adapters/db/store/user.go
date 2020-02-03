// Package store ... 
// TODO: Move into adapters/db/gateways folder
package store

import (
	"fmt"

	"github.com/Liberkeys/api-skeleton/adapters/db"
	"github.com/Liberkeys/api-skeleton/ports/models"
)

// GormUserStore ...
type GormUserStore struct {
	db *db.GormAdaptor
}

// NewGormUserStore ...
func NewGormUserStore (db *db.GormAdaptor) *GormUserStore {
	return &GormUserStore{
		db: db,
	}
}

// SelectAll ...
func (db *GormUserStore) SelectAll() ([]models.User, error) {
	fmt.Println("ADAPTER db.UsersSelectAll")

	// Do some query 

	users := []models.User{models.User{Name:"bob", Age: 123}}
	
	return users, nil
}