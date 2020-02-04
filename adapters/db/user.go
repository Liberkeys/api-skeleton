package db

import (
	"fmt"

	"github.com/Liberkeys/api-skeleton/models")

// UserStore ...
type UserStore struct {
	db *Driver
}

// NewUserStore ...
func NewUserStore(db *Driver) *UserStore {
	return &UserStore{
		db: db,
	}
}

// SelectAll ...
func (store *UserStore) SelectAll() ([]*models.User, error) {
	fmt.Println("ADAPTER db.SelectAll")

	// Do some query

	users := []*models.User{
		{
			Name: "alice",
			Age:  404,
		},
		{
			Name: "bob",
			Age:  123,
		},
	}

	return users, nil
}

// SelectOneByID ...
func (store *UserStore) SelectOneByID(id string) (*models.User, error) {
	fmt.Println("ADAPTER db.SelectOne")

	// Do some query

	user := &models.User{
		Name: "bob",
		Age:  123,
	}

	return user, nil
}

// InsertOne ...
func (store *UserStore) InsertOne() (*models.User, error) {
	return nil, nil
}

// UpdateOne ...
func (store *UserStore) UpdateOne() (*models.User, error) {
	return nil, nil
}

// ArchiveOne ...
func (store *UserStore) ArchiveOne() (*models.User, error) {
	return nil, nil
}