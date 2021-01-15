package router

import (
	"github.com/TreeHole-ccnu/TreeHole-backend/handler"
	"github.com/TreeHole-ccnu/TreeHole-backend/handler/sd"
	"github.com/TreeHole-ccnu/TreeHole-backend/middleware"
	//ginSwagger "github.com/swaggo/gin-swagger"
	//"github.com/swaggo/gin-swagger/swaggerFiles"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	// Middlewares.
	g.Use(gin.Recovery())
	g.Use(middleware.NoCache)
	g.Use(middleware.Options)
	g.Use(middleware.Secure)
	g.Use(mw...)
	// 404 Handler.
	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "The incorrect API route.")
	})
	// The health check handlers

	//g.GET("/Swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	svcd := g.Group("/sd")
	{
		svcd.GET("/health", sd.HealthCheck)
		svcd.GET("/disk", sd.DiskCheck)
		svcd.GET("/cpu", sd.CPUCheck)
		svcd.GET("/ram", sd.RAMCheck)
	}


	g.Use(middleware.JwtAAuth())

	user := g.Group("/treehole/v1/user")
	{
		user.GET("/verification", handler.SendVer) 		//发送手机验证码
		user.POST("/login",handler.UserLogin) 			//用户登录
		user.POST("/register", handler.UserRegister)		//用户注册
	}
	return g
}

