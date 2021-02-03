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

	g.POST("/treehole/v1/user/login", handler.UserLogin) //用户登录

	g.GET("/treehole/v1/user/verification", handler.SendVer) //发送用户手机验证码

	g.POST("/treehole/v1/user/register", handler.UserRegister) //用户注册

	g.POST("/treehole/v1/user/resetting", handler.UserResetting) //用户忘记密码，修改密码

	g.Use(middleware.JwtAAuth()) //身份验证

	user := g.Group("/treehole/v1/user")
	{

		user.GET("/information", handler.InfoGetting)    //获得用户信息
		user.POST("/information", handler.InfoResetting) //修改用户基本信息
		user.POST("/image", handler.UserImage)
	}

	g.GET("/treehole/v1/image", handler.Image) //上传照片

	administrator := g.Group("/treehole/v1/administrator")
	{
		administrator.POST("/addition", handler.LevelChanging)      //超级管理员添加管理员
		administrator.POST("/verification", handler.StatusChanging) //管理员通过志愿者申请
		administrator.GET("/verification", handler.Verification)    //批量查看志愿者申请
		administrator.POST("/information", handler.VolunteerSearch) //查找志愿者
		administrator.GET("/information", handler.GetDetailedInfo)  //志愿者信息详情页
	}

	volunteer := g.Group("/treehole/v1/volunteer")
	{
		volunteer.POST("/information", handler.VolunteerInfo) //志愿者申请信息
		volunteer.GET("/checking", handler.VolunteerCheck)    //获取志愿者申请进度
	}

	return g
}
