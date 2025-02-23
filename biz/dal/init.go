package dal

import "lonely-monitor/biz/dal/db"

// Init init dal
func Init() {
	db.GetDB() // mysql init
}
