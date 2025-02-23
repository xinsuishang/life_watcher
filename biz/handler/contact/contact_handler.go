package contact

import (
	"context"
	"lonely-monitor/biz/model/contact"
	service "lonely-monitor/biz/service/contact"
	"lonely-monitor/pkg/utils"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// HandleAddContact 处理添加联系人
func HandleAddContact(ctx context.Context, c *app.RequestContext) {
	var req contact.AddContactRequest
	if err := c.BindAndValidate(&req); err != nil {
		c.JSON(consts.StatusBadRequest, utils.Error(400, "无效的请求参数"))
		return
	}

	_, err := service.NewContactService(ctx, c).AddContact(&req)
	if err != nil {
		c.JSON(consts.StatusInternalServerError, utils.Error(500, "添加联系人失败"))
		return
	}

	c.JSON(consts.StatusOK, utils.Success("ok"))
}

// HandleGetContacts 处理获取联系人列表
func HandleGetContacts(ctx context.Context, c *app.RequestContext) {
	contacts, err := service.NewContactService(ctx, c).GetContacts()
	if err != nil {
		c.JSON(consts.StatusInternalServerError, utils.Error(500, "获取联系人列表失败"))
		return
	}

	c.JSON(consts.StatusOK, utils.Success(contacts))
}

// HandleDeleteContact 处理删除联系人
func HandleDeleteContact(ctx context.Context, c *app.RequestContext) {
	var req contact.DeleteContactRequest
	if err := c.BindAndValidate(&req); err != nil {
		c.JSON(consts.StatusBadRequest, utils.Error(400, "无效的请求参数"))
		return
	}

	err := service.NewContactService(ctx, c).DeleteContact(req.ContactID)
	if err != nil {
		c.JSON(consts.StatusInternalServerError, utils.Error(500, "删除联系人失败"))
		return
	}

	c.JSON(consts.StatusOK, utils.Success("ok"))
}
