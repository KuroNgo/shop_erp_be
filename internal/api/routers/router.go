package routers

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"shop_erp_mono/internal/api/routers/accounting_management"
	"shop_erp_mono/internal/api/routers/human_resources_management"
	"shop_erp_mono/internal/api/routers/log_activity"
	"shop_erp_mono/internal/api/routers/sales_and_distributing_management"
	swaggerroute "shop_erp_mono/internal/api/routers/swagger"
	"shop_erp_mono/internal/api/routers/warehouse_management"
	"shop_erp_mono/internal/config"
	cronjob "shop_erp_mono/pkg/interface/cron"
	"time"
)

func SetUp(env *config.Database, cr *cronjob.CronScheduler, timeout time.Duration, db *mongo.Database, client *mongo.Client, gin *gin.Engine, cacheTTL time.Duration) {
	swaggerroute.SwaggerRouter(env, timeout, db, gin.Group(""))
	log_activity.ActivityRoute(env, timeout, db, client, gin.Group("/api/v1"), cacheTTL)
	human_resources_management.SetUp(env, cr, timeout, db, client, gin, cacheTTL)
	accounting_management.SetUp(env, client, timeout, db, gin, cacheTTL)
	sales_and_distributing_management.SetUp(env, client, timeout, db, gin, cacheTTL)
	warehouse_management.SetUp(env, client, timeout, db, gin, cacheTTL)

	err := DataSeeds(context.Background(), client)
	if err != nil {
		fmt.Print("data seed is error")
	}

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
