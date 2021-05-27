package middleware

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model/request"
	"gin-vue-admin/model/response"
	"gin-vue-admin/service"

	"github.com/gin-gonic/gin"
)

// Interceptor
func CasbinHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, _ := c.Get("claims")
		waitUse := claims.(*request.CustomClaims)
		// Get the requested URI
		obj := c.Request.URL.RequestURI()
		// Get request method
		act := c.Request.Method
		// Get the role of the user
		sub := waitUse.AuthorityId
		e := service.Casbin()
		// Is there a strategy?
		success, _ := e.Enforce(sub, obj, act)
		if global.GVA_CONFIG.System.Env == "develop" || success {
			c.Next()
		} else {
			response.FailWithDetailed(gin.H{}, "Insufficient permissions", c)
			c.Abort()
			return
		}
	}
}
