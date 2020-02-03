package emails

import (
	"github.com/Liberkeys/api-skeleton/ports/models"
	"fmt"
)

// NotifyUser ...
func (*PostmanNotifier) NotifyUser(users []models.User) error {
	fmt.Println("PORT emails.NotifyUser")
	return nil
}