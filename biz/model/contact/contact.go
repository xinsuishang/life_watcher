package contact

// AddContactRequest 添加联系人请求
type AddContactRequest struct {
	Type        string `json:"type" binding:"required"`  // 联系方式类型
	Value       string `json:"value" binding:"required"` // 联系方式值
	IsEmergency bool   `json:"is_emergency"`             // 是否为紧急联系人
}

// GetContactsRequest 获取联系人列表请求
type GetContactsRequest struct {
	UserID int64 `json:"-" vd:"$!='userID'"` // 用户ID，从上下文获取
}

// DeleteContactRequest 删除联系人请求
type DeleteContactRequest struct {
	ContactID string `path:"id" binding:"required"` // 联系人ID，从路径参数获取
}

// Response 联系人响应
type Response struct {
	ID          int64  `json:"id"`           // 联系人ID
	Type        string `json:"type"`         // 联系方式类型
	Value       string `json:"value"`        // 联系方式值
	IsEmergency bool   `json:"is_emergency"` // 是否为紧急联系人
}
