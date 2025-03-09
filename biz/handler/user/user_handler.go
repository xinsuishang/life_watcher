package user

import (
	"context"
	"lonely-monitor/biz/model/user"
	service "lonely-monitor/biz/service/user"
	"lonely-monitor/pkg/errno"
	"lonely-monitor/pkg/utils"

	"github.com/cloudwego/hertz/pkg/app"
)

// HandleRegister 处理用户注册
func HandleRegister(c context.Context, ctx *app.RequestContext) {
	var req user.RegisterRequest
	if err := ctx.BindAndValidate(&req); err != nil {
		ctx.JSON(errno.HttpSuccess, utils.Error(ctx, errno.ParamErrCode, "无效的请求参数"))
		return
	}

	_, err := service.NewUserService(c, ctx).UserRegister(&req)
	if err != nil {
		ctx.JSON(errno.HttpSuccess, utils.Error(ctx, errno.ServiceErrCode, "注册失败"))
		return
	}

	ctx.JSON(errno.HttpSuccess, utils.Success(ctx, "ok"))
}

// HandleLogin 处理用户登录
func HandleLogin(c context.Context, ctx *app.RequestContext) {
	var req user.LoginRequest
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(errno.HttpSuccess, utils.Error(ctx, errno.ParamErrCode, "无效的请求参数"))
		return
	}

	token, err := service.NewUserService(c, ctx).UserLogin(&req)
	if err != nil {
		ctx.JSON(errno.HttpSuccess, utils.Error(ctx, errno.AuthErrCode, "用户名或密码错误"))
		return
	}

	ctx.JSON(errno.HttpSuccess, utils.Success(ctx, &utils.LoginResponse{
		Token: token,
	}))
}

// HandleCheckIn 处理用户签到
func HandleCheckIn(c context.Context, ctx *app.RequestContext) {
	userID, _ := ctx.Get("userID")

	err := service.NewUserService(c, ctx).UserCheckIn(userID.(int64))
	if err != nil {
		ctx.JSON(errno.HttpSuccess, utils.Error(ctx, errno.ServiceErrCode, "签到失败"))
		return
	}

	ctx.JSON(errno.HttpSuccess, utils.Success(ctx, "ok"))
}

// HandleLogout 处理用户登出
func HandleLogout(c context.Context, ctx *app.RequestContext) {
	ctx.JSON(errno.HttpSuccess, utils.Success(ctx, "ok"))
}
