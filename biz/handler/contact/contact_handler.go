package contact

import (
	"context"
	"lonely-monitor/biz/model/contact"
	service "lonely-monitor/biz/service/contact"
	"lonely-monitor/pkg/errno"
	"lonely-monitor/pkg/utils"

	"github.com/cloudwego/hertz/pkg/app"
)

// HandleAddContact 处理添加联系人
func HandleAddContact(c context.Context, ctx *app.RequestContext) {
	var req contact.AddContactRequest
	if err := ctx.BindAndValidate(&req); err != nil {
		ctx.JSON(errno.HttpSuccess, utils.Error(ctx, errno.ParamErrCode, "无效的请求参数"))
		return
	}

	_, err := service.NewContactService(c, ctx).AddContact(&req)
	if err != nil {
		ctx.JSON(errno.HttpSuccess, utils.Error(ctx, errno.ServiceErrCode, "添加联系人失败"))
		return
	}

	ctx.JSON(errno.HttpSuccess, utils.Success(ctx, "ok"))
}

// HandleGetContacts 处理获取联系人列表
func HandleGetContacts(c context.Context, ctx *app.RequestContext) {
	contacts, err := service.NewContactService(c, ctx).GetContacts()
	if err != nil {
		ctx.JSON(errno.HttpSuccess, utils.Error(ctx, errno.ServiceErrCode, "获取联系人列表失败"))
		return
	}

	ctx.JSON(errno.HttpSuccess, utils.Success(ctx, contacts))
}

// HandleDeleteContact 处理删除联系人
func HandleDeleteContact(c context.Context, ctx *app.RequestContext) {
	var req contact.DeleteContactRequest
	if err := ctx.BindAndValidate(&req); err != nil {
		ctx.JSON(errno.HttpSuccess, utils.Error(ctx, errno.ParamErrCode, "无效的请求参数"))
		return
	}

	err := service.NewContactService(c, ctx).DeleteContact(req.ContactID)
	if err != nil {
		ctx.JSON(errno.HttpSuccess, utils.Error(ctx, errno.ServiceErrCode, "删除联系人失败"))
		return
	}

	ctx.JSON(errno.HttpSuccess, utils.Success(ctx, "ok"))
}
