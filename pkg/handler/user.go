package handler

import (
	"wp-demo/pkg/domain/service"

	"github.com/gin-gonic/gin"
)

type RegisterReq struct {
	Username string
	Password string
}

func Register(userSrv *service.UserService) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var request RegisterReq
		err := ctx.ShouldBindJSON(&request)
		if err != nil {
			ctx.JSON(500, err.Error())
			return
		}
		err = userSrv.Register(ctx, request.Username, request.Password)
		if err != nil {
			ctx.JSON(500, err.Error())
			return
		}
		ctx.JSON(200, "创建成功")
	}
}
