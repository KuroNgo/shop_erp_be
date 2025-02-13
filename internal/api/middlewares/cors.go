package middlewares

import "github.com/gin-gonic/gin"

var (
	host1 = "http://localhost:3000"
	host2 = "http://localhost:5173"
	host3 = "https://shop-erp-fe.vercel.app"
)

func CORSPublic() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		origin := ctx.Request.Header.Get("Origin")
		if origin == host1 || origin == host2 {
			ctx.Writer.Header().Set("Access-Control-Allow-Origin", origin)
			ctx.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
			ctx.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding")
			ctx.Writer.Header().Set("Access-Control-Allow-Methods", "GET, PUT, POST, PATCH, DELETE, OPTIONS")

			if ctx.Request.Method == "OPTIONS" {
				ctx.AbortWithStatus(204)
				return
			}

			ctx.Next()
		}
	}
}

func CORSPrivate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		origin := ctx.Request.Header.Get("Origin")
		if origin == host2 || origin == host3 {
			ctx.Writer.Header().Set("Access-Control-Allow-Origin", origin)
			ctx.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
			ctx.Writer.Header().Set("Access-Control-Allow-Headers", "Authorization,Content-Type,Content,Content-Length,Accept-Encoding")
			ctx.Writer.Header().Set("Access-Control-Allow-Methods", "GET, PUT, PATCH, POST, DELETE, OPTIONS")

			if ctx.Request.Method == "OPTIONS" {
				ctx.AbortWithStatus(204)
				return
			}

			ctx.Next()
		}
	}
}

func OptionMessages(ctx *gin.Context) {
	origin := ctx.Request.Header.Get("Origin")

	if origin == host1 || origin == host2 {
		ctx.Writer.Header().Set("Access-Control-Allow-Origin", origin)
		ctx.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		ctx.Writer.Header().Set("Access-Control-Allow-Headers", "Authorization,Content-Type,Content-Length")
		ctx.Writer.Header().Set("Access-Control-Allow-Methods", "GET, PUT, PATCH, POST, DELETE, OPTIONS")

		if ctx.Request.Method == "OPTIONS" {
			ctx.AbortWithStatus(204)
			return
		}

		ctx.Next()
	}
}
