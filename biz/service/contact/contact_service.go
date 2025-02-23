package contact

import (
	"context"
	"lonely-monitor/biz/dal/db"
	"lonely-monitor/biz/model/contact"
	"lonely-monitor/pkg/utils"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
)

type ContactService struct {
	ctx context.Context
	c   *app.RequestContext
}

// NewContactService create contact service
func NewContactService(ctx context.Context, c *app.RequestContext) *ContactService {
	return &ContactService{ctx: ctx, c: c}
}

// AddContact 添加联系人
func (s *ContactService) AddContact(req *contact.AddContactRequest) (int64, error) {
	// 加密联系方式
	encryptedValue, err := utils.EncryptContact(req.Value)
	if err != nil {
		return 0, err
	}

	// 创建联系人记录
	contact := &db.ContactMethod{
		Type:           req.Type,
		EncryptedValue: encryptedValue,
		IsEmergency:    req.IsEmergency,
	}

	return db.CreateContact(contact)
}

// GetContacts 获取联系人列表
func (s *ContactService) GetContacts() ([]contact.Response, error) {
	contacts, err := db.QueryContacts()
	if err != nil {
		return nil, err
	}

	response := make([]contact.Response, 0, len(contacts))
	for _, c := range contacts {
		decryptedValue, err := utils.DecryptContact(c.EncryptedValue)
		if err != nil {
			hlog.CtxErrorf(s.ctx, "解密联系方式失败，contact_id: %d, 错误: %v", c.ID, err)
			continue
		}

		response = append(response, contact.Response{
			ID:          c.ID,
			Type:        c.Type,
			Value:       decryptedValue,
			IsEmergency: c.IsEmergency,
		})
	}

	return response, nil
}

// DeleteContact 删除联系人
func (s *ContactService) DeleteContact(contactID string) error {
	return db.DeleteContact(contactID)
}
