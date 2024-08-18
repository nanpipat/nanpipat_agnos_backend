package password

import (
	"github.com/gin-gonic/gin"
)

func NewPasswordHTTP(r *gin.Engine) {
	controller := PasswordController{}

	r.POST("/api/strong_password_steps", controller.GetStrongPasswordSteps)
}
