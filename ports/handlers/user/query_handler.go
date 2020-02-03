package user

import (
	"fmt"

	"github.com/Liberkeys/api-skeleton/adapters/db/store"
	"github.com/Liberkeys/api-skeleton/ports/emails"
	"github.com/Liberkeys/api-skeleton/ports/models"
)

// QueryHandler ...
type QueryHandler interface {
	GetAllUsers() ([]models.User, error) // Query
}

// DefaultQueryHandler ...
type DefaultQueryHandler struct {
	store    store.UserStore // Adaptor
	notifier emails.Notifier // Port
}

// NewQueryHandler ...
func NewQueryHandler(store store.UserStore, notifier emails.Notifier) *DefaultQueryHandler {
	return &DefaultQueryHandler{
		store:    store,
		notifier: notifier,
	}
}

// GetAllUsers ...
func (h *DefaultQueryHandler) GetAllUsers() ([]models.User, error) {
	fmt.Println("PORT user.GetUsers")

	// Call some private method with depencies ready
	h.checkSomethingElsePrivate()

	// Calling adapter
	users, _ := h.store.SelectAll() // TODO : check error

	// Calling port
	h.notifier.NotifyUser(users) // TODO : check error

	return users, nil
}

func (h *DefaultQueryHandler) checkSomethingElsePrivate() {
	fmt.Println("PORT users.checkSomethingElsePrivate")
}
