package controller

import (
	"email-service/model"
	"email-service/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func Ping(ctx *gin.Context) {
	log.Printf("参数：%v", ctx.Params)
	//ctx.JSON(http.StatusOK, gin.H{
	//	"message": "pong",
	//})
	ctx.String(200, "message:%v", "pong")
}

// 健康检查
func Health(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"status": "UP",
	})
}

// 发送邮件
func EmailHandler(context *gin.Context) {
	req := &model.EmailReq{}
	if err := context.BindJSON(req); err != nil {
		log.Printf("参数错误, req: %v", req)
		context.JSON(400, model.Result{
			Code:    400,
			Message: "参数错误",
		})
		return
	}
	log.Printf("req:%v\n", req)
	err := service.SendEmail(req.ToUser, req.UserName)
	if err != nil {
		log.Printf("SendEmail fail, err=%v", err)
		context.JSON(500, model.Result{
			Code:    500,
			Message: "服务器内部错误",
		})
		return
	}

	context.JSON(200, model.Result{
		Code:    200,
		Message: "邮件发送成功",
	})
}
