package douyin

import (
	"ByteDance/service"
	"ByteDance/middleware"
	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	// public directory is used to serve static resources
	r.Static("/static", "./public")

	apiRouter := r.Group("/douyin")

	// basic apis
	apiRouter.GET("/feed/", service.Feed)
	apiRouter.GET("/user/", middleware.JWTAuthMiddleware(), service.UserInfo)
	apiRouter.POST("/user/register/", service.Register)
	apiRouter.POST("/user/login/", service.Login)
	apiRouter.POST("/publish/action/", middleware.JWTAuthMiddleware(), service.Publish)
	apiRouter.GET("/publish/list/", middleware.JWTAuthMiddleware(),service.PublishList)

	// extra apis - I
	apiRouter.POST("/favorite/action/", service.FavoriteAction)
	apiRouter.GET("/favorite/list/", service.FavoriteList)
	apiRouter.POST("/comment/action/", service.CommentAction)
	apiRouter.GET("/comment/list/", service.CommentList)

	// extra apis - II
	apiRouter.POST("/relation/action/", service.RelationAction)
	apiRouter.GET("/relation/follow/list/", service.FollowList)
	apiRouter.GET("/relation/follower/list/", service.FollowerList)
}
