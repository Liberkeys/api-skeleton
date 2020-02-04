package user

import (
	"fmt"

	"github.com/Liberkeys/api-skeleton/infrastructure/application"
	"github.com/Liberkeys/api-skeleton/ports/models"
	"github.com/Liberkeys/api-skeleton/ports/notifier"
)

// GetAllUsers ...
func GetAllUsers(ctx *application.Context) ([]*models.User, error) {
	fmt.Println("PORT user.GetAllUsers")

	checkSomethingElsePrivate()

	// Calling adapter
	users, err := ctx.UserStore().SelectAll()
	if err != nil {
		return nil, err
	}

	// Calling port
	err = notifier.NotifyUser(ctx, users)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func GetOneUser(ctx *application.Context, id string) (*models.User, error) {
	fmt.Println("PORT user.GetOneUser")

	return ctx.UserStore().SelectOneByID(id)
}

func checkSomethingElsePrivate() {
	fmt.Println("PORT users.checkSomethingElsePrivate")
}
