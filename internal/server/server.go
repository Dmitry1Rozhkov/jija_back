package server

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"jija_back/internal/config"
	"jija_back/internal/handler"
	"net/http"
)

func NewServer(
	config *config.Config,
	handlers *handler.Handler,
) *Server {

	if config.IsDebug {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()
	cfg := cors.DefaultConfig()
	cfg.AllowOrigins = append(cfg.AllowOrigins, config.CorsOrigins...)
	cfg.AllowCredentials = true
	cfg.AllowHeaders = append(cfg.AllowHeaders,
		"Access-Control-Allow-Headers",
		"Access-Control-Request-Method",
		"Access-Control-Request-Headers",
		"Accept",
		"X-Requested-With",
		"Authorization")
	router.Use(cors.New(cfg))

	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusTeapot, gin.H{"code": "URL IS INVALID", "message": "URL IS INVALID"})
	})

	atmsGroup := router.Group("/atms")
	{
		atmsGroup.GET("/", handlers.GetAtmsInfo)
	}

	officesGroup := router.Group("/offices")
	{
		officesGroup.GET("/", handlers.GetOfficesInfo)
	}

	return &Server{
		config:    config,
		GinRouter: router,
	}
}

type Server struct {
	config    *config.Config
	GinRouter *gin.Engine
}

var allowOriginFunc = func(r *http.Request) bool {
	return true
}

type errorResponse struct {
	Message string `json:"message"`
}

func newErrorResponse(c *gin.Context, statusCode int, msg string) {
	c.AbortWithStatusJSON(statusCode, errorResponse{msg})
}
