package infrastructor

import (
	mongo_driven "go.mongodb.org/mongo-driver/mongo"
	"shop_erp_mono/bootstrap"
)

type Application struct {
	Env     *bootstrap.Database
	MongoDB *mongo_driven.Client
}

func App() *Application {
	app := &Application{}
	app.Env = bootstrap.NewEnv()
	app.MongoDB = NewMongoDatabase(app.Env)
	return app
}

func (app *Application) CloseDBConnection() {
	CloseMongoDBConnection(app.MongoDB)
}
