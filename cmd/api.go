package cmd

import (
	"github.com/gin-gonic/gin"
	"github.com/nanpipat/nanpipat-agnos-backend/internal/config"
	"github.com/nanpipat/nanpipat-agnos-backend/internal/core"
	"github.com/nanpipat/nanpipat-agnos-backend/internal/server"
	"github.com/nanpipat/nanpipat-agnos-backend/modules/password"
)

func APIRun() {
	cfg := config.New()
	srv := server.New(cfg)
	db := srv.DB()

	r := gin.Default()

	r.Use(func(c *gin.Context) {
		ctx := core.NewContext(c.Request.Context(), db)
		c.Set("context", ctx)
		c.Next()
	})

	r.GET("/healthz", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	password.NewPasswordHTTP(r)

	if err := r.Run(":" + cfg.Port); err != nil {
		panic(err)
	}
}
