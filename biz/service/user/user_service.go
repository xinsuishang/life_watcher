package user

import (
	"context"
	"time"

	"lonely-monitor/biz/dal/db"
	"lonely-monitor/biz/model/user"
	"lonely-monitor/pkg/errno"
	"lonely-monitor/pkg/utils"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
)

type UserService struct {
	ctx context.Context
	c   *app.RequestContext
}

// NewUserService create user service
func NewUserService(ctx context.Context, c *app.RequestContext) *UserService {
	return &UserService{ctx: ctx, c: c}
}

// UserRegister register user return user id.
func (s *UserService) UserRegister(req *user.RegisterRequest) (user_id int64, err error) {
	user, err := db.QueryUser(req.Username)
	if err != nil {
		return 0, err
	}
	if *user != (db.User{}) {
		return 0, errno.ParamErr
	}

	salt := utils.GenerateSalt()
	user = &db.User{
		Username: req.Username,
		Password: utils.HashPassword(req.Password, salt),
		Letter:   req.Letter,
		Salt:     salt,
	}

	return db.CreateUser(user)
}

// UserLogin 用户登录
func (s *UserService) UserLogin(req *user.LoginRequest) (string, error) {
	user, err := db.QueryUser(req.Username)
	if err != nil {
		return "", err
	}
	if *user == (db.User{}) {
		return "", errno.AuthorizationFailedErr
	}

	if utils.HashPassword(req.Password, user.Salt) != user.Password {
		return "", errno.AuthorizationFailedErr
	}

	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}

// UserCheckIn 用户签到
func (s *UserService) UserCheckIn(userID int64) error {
	now := time.Now()

	// 开启事务
	tx := db.GetDB().Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// 创建签到记录
	checkIn := &db.AlertRecord{
		UserID:          userID,
		LastCheckInTime: now,
		AlertTime:       now.Add(24 * time.Hour), // 延后24小时
		Status:          0,                       // 重置通知状态
	}
	if err := tx.Create(checkIn).Error; err != nil {
		tx.Rollback()
		hlog.CtxErrorf(s.ctx, "创建签到记录失败，错误: %v", err)
		return err
	}

	// 更新预警记录
	if err := tx.Model(&db.AlertRecord{}).
		Where("user_id = ?", userID).
		Updates(map[string]interface{}{
			"last_check_in_time": now,
			"alert_time":         now.Add(24 * time.Hour), // 延后24小时
			"status":             0,                       // 重置通知状态
			"updated_at":         now,
		}).Error; err != nil {
		tx.Rollback()
		hlog.CtxErrorf(s.ctx, "更新预警记录失败，错误: %v", err)
		return err
	}

	if err := tx.Commit().Error; err != nil {
		hlog.CtxErrorf(s.ctx, "提交事务失败，错误: %v", err)
		return err
	}

	return nil
}
