package config

import (
	"app_es_service_go/providers"
	"talkcheap.xiaoeknow.com/xiaoetong/eframe/library"
)

var GormConfigs []*library.GormConfig

func InitGormConfigs(env *library.Env) {
	GormConfigs = []*library.GormConfig{
		{
			Receiver:       &providers.CoreGorm,
			ConnectionName: "gorm-core",
			DBName:         env.GetString("DB_CORE_RW_NAME"),
			Host:           env.GetString("DB_CORE_RW_HOST"),
			Port:           env.GetString("DB_CORE_RW_PORT"),
			UserName:       env.GetString("DB_CORE_RW_USERNAME"),
			Password:       env.GetString("DB_CORE_RW_PASSWORD"),
			MaxLifeTime:    env.GetInt("DB_MAX_LIFE_TIME"),
			MaxOpenConn:    env.GetInt("DB_MAX_OPEN_CONN"),
			MaxIdleConn:    env.GetInt("DB_MAX_IDLE_CONN"),
		},
		{
			Receiver: &providers.SecondGorm,
			ConnectionName: "gorm-second",
			DBName:         env.GetString("DB_SUB_RW_NAME"),
			Host:           env.GetString("DB_SUB_RW_HOST"),
			Port:           env.GetString("DB_SUB_RW_PORT"),
			UserName:       env.GetString("DB_SUB_RW_USERNAME"),
			Password:       env.GetString("DB_SUB_RW_PASSWORD"),
			MaxLifeTime:    env.GetInt("DB_MAX_LIFE_TIME"),
			MaxOpenConn:    env.GetInt("DB_MAX_OPEN_CONN"),
			MaxIdleConn:    env.GetInt("DB_MAX_IDLE_CONN"),
		},
	}
}
