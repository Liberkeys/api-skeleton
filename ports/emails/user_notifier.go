package emails

import (
	"fmt"

	"github.com/Liberkeys/api-skeleton/ports/models"
)

// NotifyUser ...
func (*PostmanNotifier) NotifyUser(users []models.User) error {
	fmt.Println("PORT emails.NotifyUser")
	return nil
}
