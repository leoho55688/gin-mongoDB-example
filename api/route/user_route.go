package route

import (
	"time"

	"backend/api/controller/user"
	"backend/bootstrap"
	"backend/model/domain"
	"backend/mongo"
	"backend/repository"
	"backend/usecase"

	"github.com/gin-gonic/gin"
)

func InitUserRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, router *gin.RouterGroup) {
	userRouter := router.Group("/user")

	NewProfileRouter(env, timeout, db, userRouter)
}

func NewProfileRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, router *gin.RouterGroup) {
	ur := repository.NewUserRepository(db, domain.CollectionUser)
	pc := &user.ProfileController{
		ProfileUsecase: usecase.NewProfileUsecase(ur, timeout),
	}
	router.GET("/profile", pc.Fetch)
}
