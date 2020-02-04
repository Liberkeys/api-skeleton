package users

import (
	"fmt"

	"github.com/Liberkeys/api-skeleton/assets/notifications"
	"github.com/Liberkeys/api-skeleton/infrastructure/app"
	"github.com/Liberkeys/api-skeleton/models"
)

// CreateOneUser ..
func CreateOneUser(ctx *app.Context, id string) (*models.User, error) {
	fmt.Println("PORT user.CreateOneUser")

	return ctx.UserStore().InsertOne()
}

// UpdateOneUser ..
func UpdateOneUser(ctx *app.Context, id string) (*models.User, error) {
	fmt.Println("PORT user.UpdateOneUser")

	return ctx.UserStore().UpdateOne()
}

// DeleteOneUser ..
func DeleteOneUser(ctx *app.Context, id string) (*models.User, error) {
	fmt.Println("PORT user.DeleteOneUser")

	return ctx.UserStore().ArchiveOne()
}

// GetAllUsers ...
func GetAllUsers(ctx *app.Context) ([]*models.User, error) {
	fmt.Println("PORT user.GetAllUsers")

	checkSomethingElsePrivate(ctx)

	// Calling adapter
	users, err := ctx.UserStore().SelectAll()
	if err != nil {
		return nil, err
	}

	// Calling port
	err = notifications.NotifyUser(ctx, users)
	if err != nil {
		return nil, err
	}

	return users, nil
}

// GetOneUser ...
func GetOneUser(ctx *app.Context, id string) (*models.User, error) {
	fmt.Println("PORT user.GetOneUser")

	return ctx.UserStore().SelectOneByID(id)
}

func checkSomethingElsePrivate(ctx *app.Context) {
	fmt.Println("PORT users.checkSomethingElsePrivate")
}
