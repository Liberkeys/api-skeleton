package user

import (
	"fmt"

	"github.com/Liberkeys/api-skeleton/infrastructure/application"
	"github.com/Liberkeys/api-skeleton/ports/models"
)

func CreateOneUser(ctx *application.Context, id string) (*models.User, error) {
	fmt.Println("PORT user.CreateOneUser")

	return ctx.UserStore().InsertOne()
}

func UpdateOneUser(ctx *application.Context, id string) (*models.User, error) {
	fmt.Println("PORT user.UpdateOneUser")

	return ctx.UserStore().UpdateOne()
}

func DeleteOneUser(ctx *application.Context, id string) (*models.User, error) {
	fmt.Println("PORT user.DeleteOneUser")

	return ctx.UserStore().ArchiveOne()
}
