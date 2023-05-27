package router

import (
	"github.com/gin-gonic/gin"
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
	/*mailGroup := router.Group("/mail")
	{
		mailGroup.GET("/download", download.Download)
	}*/

	router.Static("/static/css", "./common/view/css")
	router.Static("/static/js", "./common/view/js")
	router.LoadHTMLFiles("./common/view/html/tips.html")
	router.GET("/mail/download", func(c *gin.Context) {
		c.HTML(http.StatusOK, "tips.html", gin.H{})
	})

	if err := router.Run(addr); err != nil {
		panic(err)
	}

}
