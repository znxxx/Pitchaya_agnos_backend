package router

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/znxxx/Pitchaya_agnos_backend/controllers"
	"github.com/znxxx/Pitchaya_agnos_backend/middleware"
)

func SetupRouter(db *sql.DB) *gin.Engine {
	router := gin.Default()

	router.Use(middleware.LoggingMiddleware(db))
	router.POST("/api/password_steps", controllers.HandlePasswordSubmission)

	return router
}
