package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "net/http/pprof"
	"shop_erp_mono/internal/api/routers"
	"shop_erp_mono/internal/infrastructor/mongo"
	cronjob "shop_erp_mono/pkg/interface/cron"
	"time"
)

// @title Shop E-commerce ERP
// @version 1.0
// @description This is a server for Kuro API

// @contact.name API Support
// @contact.url
// @contact.email hoaiphong01012002@gmail.com

// @host localhost:8080
// @BasePath /api/v1
func main() {

	app, client := mongo.App()
	env := app.Env
	db := app.MongoDB.Database(env.DBName)
	defer app.CloseDBConnection()

	cr := cronjob.NewCronScheduler()

	timeout := time.Duration(env.ContextTimeout) * time.Second
	cacheTTL := time.Minute * 5

	_gin := gin.Default()

	routers.SetUp(env, cr, timeout, db, client, _gin, cacheTTL)
	fmt.Println("Location Server Web of us: http://localhost:8080")
	err := _gin.Run(env.ServerAddress)
	if err != nil {
		return
	}
}
