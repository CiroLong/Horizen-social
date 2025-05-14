package main

import (
	"github.com/CiroLong/Horizon-social/internal/controller"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	// 初始化 Gin 引擎
	r := gin.Default()

	// 注册路由
	controller.RegisterUserRoutes(r)

	// TODO: use config
	// 启动服务器
	log.Println("Server starting on :8080...")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
