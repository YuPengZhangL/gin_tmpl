package middleware

import (
	"gg_web_tmpl/common/consts"
	"gg_web_tmpl/common/resp"
	"gg_web_tmpl/common/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

// AuthMW JWT鉴权中间件
func AuthMW() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Header中获取token
		token := c.GetHeader("Authorization")
		if token == "" {
			c.JSON(http.StatusUnauthorized, resp.NewBaseResp(consts.CodeTokenError, "that's not even a token"))
			c.Abort()
			return
		}
		j := utils.NewJWT()
		// 解析token包含的信息
		claims, err := j.ParseToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, resp.NewBaseResp(consts.CodeTokenError, err.Error()))
			c.Abort()
			return
		}
		// 继续交由下一个路由处理,并将解析出的信息传递下去
		c.Set("claims", claims)
		c.Set("user_id", claims.Id)
	}
}
