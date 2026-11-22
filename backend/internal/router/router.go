package router

import (
	"net/http"
	"stockflow/internal/middleware"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func NewRouter(db *sqlx.DB) *gin.Engine {
	engine := gin.New()
	engine.Use(gin.Recovery())

	engine.Use(middleware.CORSMiddleware())
	engine.Use(middleware.RequestIDMiddleware())
	engine.Use(middleware.RequestLoggerMiddleware())

	engine.GET("/healthz", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	// Example of how to set up the router with handlers, usecases, and repositories:
	// exampleRepository := repo.NewExampleRepository(db)
	// exampleUseCase := usecase.NewExampleUsecase(exampleRepository)
	// exampleHandler := NewExampleHandler(exampleUseCase)

	// application := engine.Group("/api")
	// {
	// 	application.GET("/example", exampleHandler.HandleExample)
	// }

	return engine
}
