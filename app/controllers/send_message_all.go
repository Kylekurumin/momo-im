package controllers

import (
	"fmt"
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

	fmt.Println("http_request 给全员发送了消息", body.Msg)

	ctx.JSON(http.StatusOK, nil)
}
