package route

import (
	"time"

	"backend/api/middleware"
	"backend/bootstrap"
	"backend/mongo"

	"github.com/gin-gonic/gin"
)

func Setup(env *bootstrap.Env, timeout time.Duration, db mongo.Database, gin *gin.Engine) {
	// Initial Public Router
	publicRouter := gin.Group("/api")

	// Initial Auth Router
	InitAuthRouter(env, timeout, db, publicRouter)

	// Initial Protected Router
	protectedRouter := gin.Group("/api")
	protectedRouter.Use(middleware.JwtAuthMiddleware(env.AccessTokenSecret))

	// Initial User Router
	InitUserRouter(env, timeout, db, protectedRouter)
}
