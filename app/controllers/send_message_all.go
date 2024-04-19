package controllers

import (
	"momo-im/app/schemes"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SendMessageAll(ctx *gin.Context) {
	body := &schemes.SendMessageAllRequestBody{}
	err := ctx.ShouldBindJSON(body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, nil)
}
