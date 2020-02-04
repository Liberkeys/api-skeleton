package application

import (
	"context"
	"fmt"

	"github.com/Liberkeys/api-skeleton/adapters/db"
	"github.com/Liberkeys/api-skeleton/adapters/email"
	"github.com/Liberkeys/api-skeleton/adapters/sms"
)

type ContextMode uint8

const (
	ContextModeServer = ContextMode(1)
)

type Context struct {
	ctx          context.Context
	dbAdapter    *db.Driver
	userStore    *db.UserStore
	emailAdapter *email.PostmarkAdapter
	smsAdapter   *sms.TwilioAdapter
}

func NewContext(mode ContextMode) (*Context, error) {
	ctx := &Context{
		ctx: context.Background(),
	}

	switch mode {
	case ContextModeServer:
		err := ctx.createServer()
		if err != nil {
			return nil, err
		}

	default:
		return nil, fmt.Errorf("invalid mode for application context: %d", mode)
	}

	return ctx, nil
}

func (ctx *Context) DBAdapter() *db.Driver {
	return ctx.dbAdapter
}

func (ctx *Context) createDBAdapter() *db.Driver {
	if ctx.dbAdapter != nil {
		return ctx.dbAdapter
	}
	ctx.dbAdapter = db.NewDriver("127.0.0.1", "5432", "postgres", "", "skeleton_db")
	return ctx.dbAdapter
}

func (ctx *Context) UserStore() *db.UserStore {
	return ctx.userStore
}

func (ctx *Context) createUserStore() *db.UserStore {
	if ctx.userStore != nil {
		return ctx.userStore
	}
	ctx.userStore = db.NewUserStore(ctx.createDBAdapter())
	return ctx.userStore
}

func (ctx *Context) EmailAdapter() *email.PostmarkAdapter {
	return ctx.emailAdapter
}

func (ctx *Context) createEmailAdapter() *email.PostmarkAdapter {
	if ctx.emailAdapter != nil {
		return ctx.emailAdapter
	}
	ctx.emailAdapter = email.NewPostmarkAdapter("01E07S23QY4MSS4NZMJHEYKN7T")
	return ctx.emailAdapter
}

func (ctx *Context) SmsAdapter() *sms.TwilioAdapter {
	return ctx.smsAdapter
}

func (ctx *Context) createSmsAdapter() *sms.TwilioAdapter {
	if ctx.smsAdapter != nil {
		return ctx.smsAdapter
	}
	ctx.smsAdapter = sms.NewTwilioAdapter("9972e490-6c89-4015-bca7-e609532f4962", "9972e490")
	return ctx.smsAdapter
}

func (ctx *Context) createServer() error {
	ctx.createUserStore()
	ctx.createSmsAdapter()
	ctx.createEmailAdapter()
	return nil
}
