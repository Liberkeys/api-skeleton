package notifier

import (
	"fmt"

	"github.com/Liberkeys/api-skeleton/infrastructure/application"
	"github.com/Liberkeys/api-skeleton/ports/models"
)

// NotifyUser ...
func NotifyUser(ctx *application.Context, users []*models.User) error {
	fmt.Println("PORT emails.NotifyUser")

	for i := range users {
		err := ctx.EmailAdapter().Send(fmt.Sprintf("hello %s!", users[i].Name))
		if err != nil {
			// TODO Log error
			_ = err
		}
		err = ctx.SmsAdapter().Send("0", "1", fmt.Sprintf("hello %s!", users[i].Name))
		if err != nil {
			// TODO Log error
			_ = err
		}
	}

	return nil
}
