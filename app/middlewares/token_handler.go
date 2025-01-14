package middlewares

import (
	"github.com/acheong08/DuckDuckGo-API/app/config"
	"github.com/gin-gonic/gin"
	"strings"
)

// TokenJWTAuth 拦截器
func TokenJWTAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		// 如果没配置，则直接通过
		if len(config.LoadConfig().AccessToken) == 0 {
			ctx.Next()
			return
		}

		// 这里做简单校验
		token := ctx.Request.Header.Get("Authorization")
		token = strings.Replace(token, "Bearer ", "", -1)
		if token != config.LoadConfig().AccessToken {
			ResponseJson(ctx, 403, "非法访问", nil)
			ctx.Abort()
			return
		}

		ctx.Next()
	}

}

func ResponseJson(ctx *gin.Context, code int, errorMsg string, data interface{}) {
	ctx.JSON(code, gin.H{
		"code":     code,
		"errorMsg": errorMsg,
		"data":     data,
	})
	ctx.Abort()
}
