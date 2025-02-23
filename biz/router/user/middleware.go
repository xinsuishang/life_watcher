package user

import (
	"github.com/cloudwego/hertz/pkg/app"
	"lonely-monitor/biz/mw/jwt"
)

func rootMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _userMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _user0Mw() []app.HandlerFunc {
	// your code...
	return []app.HandlerFunc{
		jwt.JWTAuth(),
	}
}

func _loginMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _userloginMw() []app.HandlerFunc {
	// your code...
	return []app.HandlerFunc{
		jwt.JWTAuth(),
	}
}

func _registerMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _userregisterMw() []app.HandlerFunc {
	// your code...
	return nil
}
