package api

import (
	"github.com/gin-gonic/gin"
	"github.com/hq0101/go-clamav/docs"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}

func Run() error {
	docs.SwaggerInfo.Title = "ClamAV API"
	docs.SwaggerInfo.Description = "This is a sample server for ClamAV"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	r := gin.Default()
	// CORS configuration
	r.Use(CORSMiddleware())
	r.Use(gin.Recovery())

	r.GET("/healthz", func(c *gin.Context) { c.Status(http.StatusOK) })
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))

	clamd := NewClamd()

	r.GET("/ping", clamd.Ping)
	r.GET("/version", clamd.Version)
	r.GET("/versioncommands", clamd.Version)
	r.GET("/stats", clamd.Stats)
	r.POST("/reload", clamd.Reload)
	r.POST("/shutdown", clamd.Shutdown)
	r.GET("/scan", clamd.Scan)
	r.GET("/contscan", clamd.Contscan)
	r.GET("/multiscan", clamd.MultiScan)
	r.GET("/allmatchscan", clamd.MultiScan)
	r.POST("/instream", clamd.Instream)

	// Use ginSwagger middleware to serve the API docs
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r.Run(GetCfg().Listen)
}
