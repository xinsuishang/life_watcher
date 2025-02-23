package user

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	user "lonely-monitor/biz/handler/user"
)

func Register(r *server.Hertz) {

	// 用户相关路由
	root := r.Group("/", rootMw()...)
	{
		_user := root.Group("/api/user", _userMw()...)
		{

			{
				_register := _user.Group("/register", _registerMw()...)
				_register.POST("/", append(_userregisterMw(), user.HandleRegister)...)

			}
			{
				_login := _user.Group("/login", _loginMw()...)
				_login.POST("/", append(_userloginMw(), user.HandleLogin)...)

			}
		}
	}
}
