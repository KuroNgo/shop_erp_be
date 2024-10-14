package infrastructor

import (
	mongodriven "go.mongodb.org/mongo-driver/mongo"
	"shop_erp_mono/bootstrap"
)

type Application struct {
	Env     *bootstrap.Database
	MongoDB *mongodriven.Client
}

func App() (*Application, *mongodriven.Client) {
	app := &Application{}
	app.Env = bootstrap.NewEnv()
	app.MongoDB = NewMongoDatabase(app.Env)
	return app, app.MongoDB
}

func (app *Application) CloseDBConnection() {
	CloseMongoDBConnection(app.MongoDB)
}
