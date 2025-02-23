package contact

import (
	"lonely-monitor/biz/handler/contact"
	"lonely-monitor/biz/mw/jwt"

	"github.com/cloudwego/hertz/pkg/app/server"
)

// Register 注册联系人相关路由
func Register(r *server.Hertz) {
	// 联系人相关路由
	root := r.Group("/api/v1")
	{
		contacts := root.Group("/contacts", jwt.JWTAuth())
		{
			// 创建联系人
			contacts.POST("", append(createContactMw(), contact.HandleAddContact)...)
			// 获取联系人列表
			contacts.GET("", append(listContactsMw(), contact.HandleGetContacts)...)
			// 删除联系人
			contacts.DELETE("/:id", append(deleteContactMw(), contact.HandleDeleteContact)...)
		}
	}
}
