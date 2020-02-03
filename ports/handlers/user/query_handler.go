package user

import (
	"github.com/Liberkeys/api-skeleton/adapters/db/store"
	"github.com/Liberkeys/api-skeleton/ports/emails"
	"github.com/Liberkeys/api-skeleton/ports/models"
	"fmt"
)

// QueryHandler ...
type QueryHandler interface {
	GetAllUsers() ([]models.User, error) // Query
}

// DefaultQueryHandler ...
type DefaultQueryHandler struct {
	Store store.UserStore // Adaptor
	Notifier emails.Notifier // Port
}

// GetAllUsers ...
func (h *DefaultQueryHandler) GetAllUsers() ([]models.User, error) {
	fmt.Println("PORT user.GetUsers")

	// Call some private method with depencies ready
	h.checkSomethingElsePrivate()

	// Calling adapter
	users, _ := h.Store.SelectAll() // TODO : check error

	// Calling port
	h.Notifier.NotifyUser(users) // TODO : check error

	return users, nil
}

func(h *DefaultQueryHandler) checkSomethingElsePrivate() {
	fmt.Println("PORT users.checkSomethingElsePrivate")
}