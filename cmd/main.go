package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "net/http/pprof"
	"shop_erp_mono/api/routers"
	"shop_erp_mono/infrastructor"
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

	app := infrastructor.App()

	env := app.Env

	db := app.MongoDB.Database(env.DBName)
	defer app.CloseDBConnection()

	timeout := time.Duration(env.ContextTimeout) * time.Second

	_gin := gin.Default()

	routers.SetUp(env, timeout, db, _gin)
	fmt.Println("Location Server Web of us: http://localhost:8080")
	err := _gin.Run(env.ServerAddress)
	if err != nil {
		return
	}

}
