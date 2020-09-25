package controller

import (
	"email-service/model"
	"email-service/service"
	"encoding/json"
	"fmt"
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
func EmailHandler(ctx *gin.Context) {
	req := &model.EmailReq{}
	if err := ctx.BindJSON(req); err != nil {
		log.Printf("参数错误, req: %v", req)
		ctx.JSON(400, model.Result{
			Code:    400,
			Message: "参数错误",
		})
		return
	}
	log.Printf("req:%v\n", req)
	err := service.SendEmail(req.ToUser, req.UserName)
	if err != nil {
		log.Printf("SendEmail fail, err=%v", err)
		ctx.JSON(500, model.Result{
			Code:    500,
			Message: "服务器内部错误",
		})
		return
	}

	ctx.JSON(200, model.Result{
		Code:    200,
		Message: "邮件发送成功",
	})
}

func UserHandler(ctx *gin.Context) {
	instance := ctx.Param("instance")
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": "参数错误",
		})
		return
	}

	client := http.DefaultClient
	url := fmt.Sprintf("http://127.0.0.1:6666/microservice-user/%s?id=%s", instance, id)
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if request == nil || err != nil {
		log.Printf("new request fail, err:%v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": "服务器内部错误",
		})
		return
	}
	authorization := ctx.GetHeader("Authorization")
	log.Printf("Authorization:%v", authorization)
	request.Header.Add("Authorization", authorization)
	resp, err := client.Do(request)
	if err != nil {
		log.Printf("err: %v, resp: %v\n", err, resp)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": "服务器内部错误",
		})
		return
	}
	log.Printf("resp:%v\n\n", resp)
	var result = model.Result{}
	if 200 == resp.StatusCode {
		err = json.NewDecoder(resp.Body).Decode(&result)
		if err != nil {
			log.Printf("Unmarshal err: %v\n", err)
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"msg": "服务器内部错误",
			})
			return
		}
	}
	ctx.JSON(http.StatusOK, result)
}
