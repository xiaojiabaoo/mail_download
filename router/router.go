package router

import (
	"github.com/gin-gonic/gin"
	"mail_download/controller/cll"
	"mail_download/controller/download"
	"mail_download/controller/system"
	"net/http"
)

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")
		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		// 处理请求
		c.Next()
	}
}

func Run(addr string) {
	router := gin.Default()
	router.Use(Cors())
	mailGroup := router.Group("/mail")
	{
		mailGroup.POST("/download", download.Download)
		mailGroup.POST("/ccl", ccl.CCL)
	}

	systemGroup := router.Group("/system")
	{
		systemGroup.GET("/list", system.List)
	}
	router.Static("/static/css", "./common/view/css")
	router.Static("/static/js", "./common/view/js")
	router.Static("/static/image", "./common/view/image")
	router.Static("/static/html", "./common/view/html")
	router.LoadHTMLFiles("./common/view/html/tips.html", "./common/view/html/catalogue.html",
		"./common/view/html/tab.html", "./common/view/html/cll.html")
	router.GET("/download", func(c *gin.Context) {
		c.HTML(http.StatusOK, "tips.html", gin.H{})
	})
	router.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "tab.html", gin.H{})
	})
	router.GET("/cll", func(c *gin.Context) {
		c.HTML(http.StatusOK, "cll.html", gin.H{})
	})
	router.GET("/catalogue", func(c *gin.Context) {
		c.HTML(http.StatusOK, "catalogue.html", gin.H{})
	})
	if err := router.Run(addr); err != nil {
		panic(err)
	}
}
