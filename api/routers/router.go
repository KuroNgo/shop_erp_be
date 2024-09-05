package routers

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"shop_erp_mono/api/routers/accounting_management"
	"shop_erp_mono/api/routers/human_resources_management"
	swaggerroute "shop_erp_mono/api/routers/swagger"
	"shop_erp_mono/bootstrap"
	"time"
)

func SetUp(env *bootstrap.Database, timeout time.Duration, db *mongo.Database, gin *gin.Engine) {
	swaggerroute.SwaggerRouter(env, timeout, db, gin.Group(""))
	human_resources_management.SetUp(env, timeout, db, gin)
	accounting_management.SetUp(env, timeout, db, gin)
}
