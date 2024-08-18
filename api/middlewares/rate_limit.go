package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sync"
	"time"
)

const (
	maxRequests     = 5
	perMinutePeriod = 15 * time.Second
)

var (
	ipRequestsCounts = make(map[string]int) // can use some distrubuted db
	mutex            = &sync.Mutex{}
)

func RateLimiter() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ip := ctx.ClientIP()
		mutex.Lock()
		defer mutex.Unlock()
		count := ipRequestsCounts[ip]
		if count >= maxRequests {
			ctx.JSON(http.StatusTooManyRequests, gin.H{
				"message": "Gửi quá nhiều request, vui lòng đợi 15 giây để thực hiện request tiếp theo!",
				"status":  "fail",
			})
			return
		}

		ipRequestsCounts[ip] = count + 1
		time.AfterFunc(perMinutePeriod, func() {
			mutex.Lock()
			defer mutex.Unlock()

			ipRequestsCounts[ip] = ipRequestsCounts[ip] - 1
		})

		ctx.Next()
	}

}
