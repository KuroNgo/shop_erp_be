package routers

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"shop_erp_mono/bootstrap"
	"time"
)

func SetUp(env *bootstrap.Database, timeout time.Duration, db *mongo.Database, gin *gin.Engine) {

}
