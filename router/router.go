package router

import (
	"github.com/gin-gonic/gin"
	"github.com/sridhar-gowda/go-gin-mysql/controller"
)

func SetupRouter() *gin.Engine {
	g := gin.Default()
	grp1 := g.Group("/api")
	{
		grp1.POST("customer", controller.CreateUser)
		grp1.POST("request", controller.PostRequest)
		grp1.GET("stats", controller.GetStats)
		grp1.POST("ipblacklist", controller.AddIPBlackList)
		grp1.PUT("customer/:id", controller.UpdateCustomerStatus)
		grp1.POST("userblacklist", controller.AddBlackListUser)
		grp1.DELETE("ipblacklist", controller.DeleteIPBlackList)

	}
	return g
}
