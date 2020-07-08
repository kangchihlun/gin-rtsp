package server

import (
	"ginrtsp/api"
	"net/http"

	"github.com/gin-gonic/gin"
)

// NewRouter Gin 路由配置
func NewRouter() *gin.Engine {
	r := gin.Default()
	r.LoadHTMLGlob("html/*.html")
	r.Static("/assets", "./html")
	// 路由
	r.GET("/ping", api.Ping)
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "view-stream.html", nil)
	})
	route := r.Group("/stream")
	{
		route.POST("/play", api.PlayRTSP)
		route.POST("/upload/:channel", api.Mpeg1Video)
		route.GET("/live/:channel", api.Wsplay)
	}

	return r
}
