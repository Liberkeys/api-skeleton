package notifications

import (
	"fmt"

	"github.com/Liberkeys/api-skeleton/infrastructure/app"
	"github.com/Liberkeys/api-skeleton/models")

// NotifyUser ...
func NotifyUser(ctx *app.Context, users []*models.User) error {
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
