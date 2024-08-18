package password

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nanpipat/nanpipat-agnos-backend/internal/core"
	"github.com/nanpipat/nanpipat-agnos-backend/requests"
	"github.com/nanpipat/nanpipat-agnos-backend/services"
)

type PasswordController struct {
}

func (c *PasswordController) GetStrongPasswordSteps(ctx *gin.Context) {
	var req requests.PasswordRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	appCtx := ctx.MustGet("context").(core.IContext)
	pwdSvc := services.NewPasswordService(appCtx)

	result, err := pwdSvc.LogPasswordCheck(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Failed to log request"})
		return
	}

	ctx.JSON(http.StatusOK, result)
}
