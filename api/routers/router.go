package routers

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"shop_erp_mono/api/routers/human_resources_management"
	swaggerroute "shop_erp_mono/api/routers/swagger"
	"shop_erp_mono/bootstrap"
	"time"
)

func SetUp(env *bootstrap.Database, timeout time.Duration, db *mongo.Database, gin *gin.Engine) {
	SwaggerRouter := gin.Group("")

	swaggerroute.SwaggerRouter(env, timeout, db, SwaggerRouter)
	human_resources_management.SetUp(env, timeout, db, gin)
}
