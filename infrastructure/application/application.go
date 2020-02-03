package application

import (
	"github.com/Liberkeys/api-skeleton/adapters/db"
	"github.com/Liberkeys/api-skeleton/adapters/db/store"
	"github.com/Liberkeys/api-skeleton/adapters/http"
	"github.com/Liberkeys/api-skeleton/adapters/http/controllers"
	"github.com/Liberkeys/api-skeleton/ports/emails"
	"github.com/Liberkeys/api-skeleton/ports/handlers/user"
)

type Application struct {
	dbAdaptor *db.GormAdaptor

	userStore store.UserStore

	notifier emails.Notifier

	userQueryHandler user.QueryHandler

	userController *controllers.UserController

	server *http.Server
}

func New() *Application {
	return &Application{}
}

func (app *Application) DBGormAdaptor() *db.GormAdaptor {
	if app.dbAdaptor != nil {
		return app.dbAdaptor
	}
	app.dbAdaptor = db.NewGormAdaptor()
	return app.dbAdaptor
}

func (app *Application) UserStore() store.UserStore {
	if app.userStore != nil {
		return app.userStore
	}
	app.userStore = store.NewGormUserStore(app.DBGormAdaptor())
	return app.userStore
}

func (app *Application) Notifier() emails.Notifier {
	if app.notifier != nil {
		return app.notifier
	}
	app.notifier = emails.NewPostmanNotifier()
	return app.notifier
}

func (app *Application) UserQueryHandler() user.QueryHandler {
	if app.userQueryHandler != nil {
		return app.userQueryHandler
	}
	app.userQueryHandler = &user.DefaultQueryHandler{ // TODO: use constructor
		Store:    app.UserStore(),
		Notifier: app.Notifier(),
	}
	return app.userQueryHandler
}

func (app *Application) UserController() *controllers.UserController {
	if app.userController != nil {
		return app.userController
	}
	/*  HTTP Adapter dependencies setup */
	app.userController = &controllers.UserController{ // TODO: use constructor
		QueryHandler: app.UserQueryHandler(),
	}
	return app.userController
}

func (app *Application) Server() *http.Server {
	if app.server != nil {
		return app.server
	}

	app.server = http.NewServer(app.UserController())
	return app.server
}
