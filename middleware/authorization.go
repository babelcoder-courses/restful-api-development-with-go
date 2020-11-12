package middleware

import (
	"course-go/models"
	"net/http"

	"github.com/casbin/casbin"
	"github.com/gin-gonic/gin"
)

func Authorize() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		auth, ok := ctx.Get("sub")
		if !ok {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		enforcer := casbin.NewEnforcer("config/acl_model.conf", "config/policy.csv")
		ok = enforcer.Enforce(auth.(*models.Auth), ctx.Request.URL.Path, ctx.Request.Method)

		if !ok {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"error": "you are not allowed to access this resource",
			})
			return
		}

		ctx.Next()
	}
}
