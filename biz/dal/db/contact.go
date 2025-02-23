// AIGC START
package db

import (
	"errors"
)

// ContactMethod 联系方式模型
type ContactMethod struct {
	BaseModel
	Type           string `json:"type" gorm:"column:type;type:varchar(255)"`                       // 联系方式类型
	EncryptedValue string `json:"encrypted_value" gorm:"column:encrypted_value;type:varchar(255)"` // 加密后的联系方式值
	IsEmergency    bool   `json:"is_emergency" gorm:"column:is_emergency;type:boolean"`            // 是否为紧急联系人
}

// CreateContact 创建联系人
func CreateContact(contact *ContactMethod) (int64, error) {
	err := GetDB().Create(contact).Error
	if err != nil {
		return 0, err
	}
	return contact.ID, nil
}

// QueryContacts 查询所有联系人
func QueryContacts() ([]ContactMethod, error) {
	var contacts []ContactMethod
	err := GetDB().Find(&contacts).Error
	if err != nil {
		return nil, err
	}
	return contacts, nil
}

// DeleteContact 删除联系人
func DeleteContact(contactID string) error {
	result := GetDB().Where("id = ?", contactID).Delete(&ContactMethod{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("联系人不存在")
	}
	return nil
}

// AIGC END
