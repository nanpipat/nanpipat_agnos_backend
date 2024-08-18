package password

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewPasswordHTTP(r *gin.Engine, db *gorm.DB) {
	controller := PasswordController{}

	r.POST("/api/strong_password_steps", controller.GetStrongPasswordSteps)
}
