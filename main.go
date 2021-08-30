package main

import (
	"github.com/gin-gonic/gin"
	"github.com/yongjie0203/go-trade-user/database"
	"github.com/yongjie0203/go-trade-user/handler"
	"github.com/yongjie0203/go-trade-user/middleware"
	"github.com/yongjie0203/go-universal/db"
	"github.com/yongjie0203/go-universal/rcache"
	"github.com/yongjie0203/go-universal/umiddleware"
)

func main() {

	rcache.SetupRedis()
	db.InitDB()
	database.TableUpdateRegister()
	db.DBUpdate()
	//go websockets.SetUpWebSocket()

	ginRouter := gin.Default()
	ginRouter.Use(middleware.Cros())
	ginRouter.Use(umiddleware.LogAopReq())

	api := ginRouter.Group("/v1")
	user := api.Group("/user")
	{
		user.POST("/register", handler.Register)
		//trade.GET("/order", handler.Order)
		//trade.POST("/list/:size/:page", controller.GetArticleList)
		//trade.GET("/get/:id", controller.GetArticle)
		//trade.GET("/delete/:id", controller.eteArticle)
	}

	ginRouter.Run()

}
