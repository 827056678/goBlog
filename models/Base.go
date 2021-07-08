package models

import (
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/server/web"
)

type BaseModel struct {
}

func init() {
	var (
		sql_type, _  = web.AppConfig.String("SQLTYPE")
		host_name, _ = web.AppConfig.String("HOSTNAME")
		data_base, _ = web.AppConfig.String("DATABASE")
		user_name, _ = web.AppConfig.String("USERNAME")
		password, _  = web.AppConfig.String("PASSWORD")
		host_port, _ = web.AppConfig.String("HOSTPORT")
		// prefix, _	= web.AppConfig.String("PREFIX")
		charset, _ = web.AppConfig.String("CHARSET")
	)
	sql_source := user_name + ":" + password + "@tcp(" + host_name + ":" + host_port + ")/" + data_base + "?charset=" + charset
	// 注册驱动
	orm.RegisterDriver(sql_type, orm.DRMySQL)
	orm.RegisterDataBase("default", sql_type, sql_source)
}
