package main

import (
	"log"

	"wp-demo/pkg/domain/service"
	"wp-demo/pkg/handler"
	"wp-demo/pkg/infrastructure/db"
	"wp-demo/pkg/infrastructure/repository"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := db.NewDB()
	if err != nil {
		panic(err)
	}

	userRepo := repository.NewUserRepository(db)
	userSrv := service.NewUserService(userRepo)

	articleRepo := repository.NewArticleRepository(db)
	articleSrv := service.NewArticleService(articleRepo)

	// Create a Gin router with default middleware (logger and recovery)
	r := gin.Default()

	// Define a simple GET endpoint
	r.POST("/register", handler.Register(userSrv))

	r.POST("/article/create", handler.CreateArticle(articleSrv))       // 创建
	r.GET("/articles/:id", handler.GetArticle(articleSrv))       // 获取单个
	r.GET("/articles", handler.ListArticle(articleSrv))          // 列表
	r.DELETE("/articles/:id", handler.DeleteArticle(articleSrv)) // 删除
	// r.PUT("/articles/:id", handler.UpdateArticle(articleSrv))    // 更新

	// Start server on port 8080 (default)
	// Server will listen on 0.0.0.0:8080 (localhost:8080 on Windows)
	if err := r.Run(); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
