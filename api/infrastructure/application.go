package infrastructure

import "go.uber.org/zap"

type Application struct {
	Env    *Env
	Db     *Orm
	Logger *zap.Logger
}

func App(envPath string) *Application {

	app := &Application{}

	app.Logger = NewLogger()

	app.Env = NewEnv(envPath, app.Logger)

	app.Db = NewORM(app.Env)

	return app
}
