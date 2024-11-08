package routers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"shop_erp_mono/internal/api/routers/accounting_management"
	"shop_erp_mono/internal/api/routers/human_resources_management"
	"shop_erp_mono/internal/api/routers/sales_and_distributing_management"
	swaggerroute "shop_erp_mono/internal/api/routers/swagger"
	"shop_erp_mono/internal/api/routers/warehouse_management"
	"shop_erp_mono/internal/config"
	"time"
)

func SetUp(env *config.Database, timeout time.Duration, db *mongo.Database, client *mongo.Client, gin *gin.Engine, cacheTTL time.Duration) {
	swaggerroute.SwaggerRouter(env, timeout, db, gin.Group(""))
	human_resources_management.SetUp(env, timeout, db, client, gin, cacheTTL)
	accounting_management.SetUp(env, timeout, db, gin)
	sales_and_distributing_management.SetUp(env, timeout, db, gin, cacheTTL)
	warehouse_management.SetUp(env, timeout, db, gin, cacheTTL)

	// Đếm các route
	routeCount := countRoutes(gin)
	fmt.Printf("The number of API endpoints: %d\n", routeCount)
}

func countRoutes(r *gin.Engine) int {
	count := 0
	routes := r.Routes()
	for range routes {
		count++
	}
	return count
}
