package db

type User struct {
	BaseModel
	Username string `json:"username" gorm:"unique,column:username;type:varchar(50)"` // 用户名
	Password string `json:"password" gorm:"column:password;type:varchar(255)"`       // 密码
	Salt     string `json:"salt" gorm:"column:salt;type:varchar(255)"`               // 盐值
	Letter   string `json:"letter" gorm:"column:letter;type:varchar(255)"`           // 信件
}

// CreateUser create user info
func CreateUser(user *User) (int64, error) {
	err := GetDB().Create(user).Error
	if err != nil {
		return 0, err
	}
	return user.ID, err
}

// QueryUser query User by user_name
func QueryUser(userName string) (*User, error) {
	var user User
	if err := GetDB().Where("username = ?", userName).Find(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// GetUserByID 根据ID获取用户信息
func GetUserByID(id int64) (*User, error) {
	var user User
	if err := GetDB().Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
