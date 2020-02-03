package emails

import (
	"github.com/Liberkeys/api-skeleton/ports/models"
)

// Notifier ...
type Notifier interface {
	NotifyUser(users []models.User) error
}
