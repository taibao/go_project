package config

import (
	"talkcheap.xiaoeknow.com/xiaoetong/eframe/library"
)

var CoreDBConfig *library.MysqlConfig
var ComponentDBConfig *library.MysqlConfig

func InitMysqlConfig(env *library.Env) {
	CoreDBConfig = &library.MysqlConfig{
		ConnectionName: "mysql-core",
		DBName:      env.GetString("DB_CORE_RW_NAME"),
		Host:        env.GetString("DB_CORE_RW_HOST"),
		Port:        env.GetString("DB_CORE_RW_PORT"),
		UserName:    env.GetString("DB_CORE_RW_USERNAME"),
		Password:    env.GetString("DB_CORE_RW_PASSWORD"),
		MaxLifeTime: env.GetInt("DB_MAX_LIFE_TIME"),
		MaxOpenConn: env.GetInt("DB_MAX_OPEN_CONN"),
		MaxIdleConn: env.GetInt("DB_MAX_IDLE_CONN"),
	}

	ComponentDBConfig = &library.MysqlConfig{
		ConnectionName: "mysql-component",
		DBName: env.GetString("DB_SUB_RW_NAME"),
		Host:        env.GetString("DB_SUB_RW_HOST"),
		Port:        env.GetString("DB_SUB_RW_PORT"),
		UserName:    env.GetString("DB_SUB_RW_USERNAME"),
		Password:    env.GetString("DB_SUB_RW_PASSWORD"),
		MaxLifeTime: env.GetInt("DB_MAX_LIFE_TIME"),
		MaxOpenConn: env.GetInt("DB_MAX_OPEN_CONN"),
		MaxIdleConn: env.GetInt("DB_MAX_IDLE_CONN"),
	}
}

