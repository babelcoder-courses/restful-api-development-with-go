package middleware

import (
	"course-go/models"
	"net/http"

	"github.com/casbin/casbin"
	"github.com/gin-gonic/gin"
)

func Authorize() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		auth, ok := ctx.Get("auth")
		if !ok {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		enforcer := casbin.NewEnforcer("config/acl_model.conf", "config/policy.csv")
		var role string
		switch auth.(*models.Auth).Role {
		case models.RoleAdmin:
			role = "Admin"
		case models.RoleEditor:
			role = "Editor"
		case models.RoleMember:
			role = "Member"
		}

		ok = enforcer.Enforce(role, ctx.Request.URL.Path, ctx.Request.Method)
		if !ok {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"error": "you are not allowed to access this resource",
			})
			return
		}

		ctx.Next()
	}
}
