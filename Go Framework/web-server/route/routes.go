package route

import (
	"database/sql"
	"web-server/handler"
	"web-server/middleware"
	auth_repo "web-server/repo/auth/auth_impl"
	analyzer_service "web-server/service/analyzer/analyzer_impl"
	auth_service "web-server/service/auth/auth_impl"

	"github.com/gin-gonic/gin"
)


func RegisterRoute(router *gin.Engine, db *sql.DB) {


     
  authRepo := auth_repo.NewAuthImpl(db)
  authSvc := auth_service.NewAuth(authRepo)

  analyzerSvc := analyzer_service.NewAnalyzerImpl(db)

  h := handler.NewHandler(authSvc, analyzerSvc)

  

  // Authentication
  router.POST("/signup", h.SignUP)
  router.POST("/login", h.Login)

  // Access After Validation
  router.POST("/count", middleware.AuthenticationMiddleware, h.Filecount)
  router.GET("/results", middleware.AuthenticationMiddleware, h.Getdata)
  router.GET("/result/:Id", middleware.AuthenticationMiddleware, h.Getid)
  router.DELETE("/result/:Id", middleware.AuthenticationMiddleware, h.Deleteid)
  router.PATCH("/result/:Id", middleware.AuthenticationMiddleware, h.Updateid)
}
