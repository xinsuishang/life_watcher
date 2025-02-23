package user

import (
	"context"
	"lonely-monitor/biz/model/user"
	service "lonely-monitor/biz/service/user"
	"lonely-monitor/pkg/utils"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// HandleRegister 处理用户注册
func HandleRegister(ctx context.Context, c *app.RequestContext) {
	var req user.RegisterRequest
	if err := c.BindAndValidate(&req); err != nil {
		c.JSON(consts.StatusBadRequest, utils.Error(400, "无效的请求参数"))
		return
	}

	_, err := service.NewUserService(ctx, c).UserRegister(&req)
	if err != nil {
		c.JSON(consts.StatusInternalServerError, utils.Error(500, "注册失败"))
		return
	}

	c.JSON(consts.StatusOK, utils.Success("ok"))
}

// HandleLogin 处理用户登录
func HandleLogin(ctx context.Context, c *app.RequestContext) {
	var req user.LoginRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(consts.StatusBadRequest, utils.Error(400, "无效的请求参数"))
		return
	}

	token, err := service.NewUserService(ctx, c).UserLogin(&req)
	if err != nil {
		c.JSON(consts.StatusUnauthorized, utils.Error(401, "用户名或密码错误"))
		return
	}

	c.JSON(consts.StatusOK, utils.Success(&utils.LoginResponse{
		Token: token,
	}))
}

// HandleCheckIn 处理用户签到
func HandleCheckIn(ctx context.Context, c *app.RequestContext) {
	userID, _ := c.Get("userID")

	err := service.NewUserService(ctx, c).UserCheckIn(userID.(int64))
	if err != nil {
		c.JSON(consts.StatusInternalServerError, utils.Error(500, "签到失败"))
		return
	}

	c.JSON(consts.StatusOK, utils.Success("ok"))
}

// HandleLogout 处理用户登出
func HandleLogout(ctx context.Context, c *app.RequestContext) {
	c.JSON(consts.StatusOK, utils.Success("ok"))
}
