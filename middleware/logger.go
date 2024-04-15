package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func RequestLogger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		t := time.Now()

		ctx.Next()

		latency := time.Since(t)

		fmt.Printf("%s %s %s %s %s\n",
			ctx.Request.Method,
			ctx.Request.RequestURI,
			ctx.Request.Proto,
			latency,
			ctx.Request.Host,
		)

	}
}

func ResponseLogger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Writer.Header().Set("X-Content-Type-Options", "nosniff")

		ctx.Next()

		fmt.Printf("%d %s %s %s\n",
			ctx.Writer.Status(),
			ctx.Request.Method,
			ctx.Request.RequestURI,
			ctx.Request.Host,
		)
	}
}
