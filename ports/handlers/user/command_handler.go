package user

import (
	"github.com/Liberkeys/api-skeleton/adapters/db"
)

// ICommandHandlers ...
type ICommandHandlers interface {
	UpdateUser(userID int) error // Command
}

// CommandHandler ...
type CommandHandler struct {
	DatabaseAdaptor db.Adaptor // Adaptor
}

// UpdateUser ...
func (handler *CommandHandler) UpdateUser(userID int) error {

	// Do something

	return nil
}
