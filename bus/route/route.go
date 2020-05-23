package route

import (
	"oceanEngineService/bus/route/middleware"

	"net/http"

	"github.com/gin-gonic/gin"
	"oceanEngineService/bus/route/handler"
)

const (
	IsAdmin = int(1)
	IsUser  = int(2)
)

func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	g.Use(gin.Recovery())
	g.Use(middleware.NoCache)
	g.Use(middleware.Options)
	g.Use(middleware.Secure)
	g.Use(mw...)

	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "The Incorrect API Route")
	})

	check := g.Group("/friendcycle/api/check")
	{
		check.GET("/health", handler.HealthCheck)
		check.GET("/disk", handler.DiskCheck)
		check.GET("/cpu", handler.CPUCheck)
		check.GET("/ram", handler.RAMCheck)
	}

	return g
}
