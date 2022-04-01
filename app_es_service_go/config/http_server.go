package config

import (
	"talkcheap.xiaoeknow.com/xiaoetong/eframe/library"
)

var HttpServerConfig *library.HttpServerConfig

func InitHttpServerConfig(env *library.Env) {
	HttpServerConfig = &library.HttpServerConfig{
		Port: env.GetInt("HTTP_SERVER_PORT"),
	}
}
