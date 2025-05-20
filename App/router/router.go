package router

import (
	"Golunwen2/App/controller"
	"Golunwen2/App/middleware"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitGin() {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"}, // 允许的前端地址
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "hello!!!!!"})
	})
	r.POST("/register", controller.Register)
	r.POST("/login", controller.Login)
	r.GET("/info", middleware.AuthMiddleware(), controller.Info)

	if err := r.Run(":8080"); err != nil {
		panic("启动失败, error=" + err.Error())
	}
}
