package config

import (
	"app_es_service_go/providers"
	"talkcheap.xiaoeknow.com/xiaoetong/eframe/library"
)

var LoggerConfigs []*library.LoggerConfig

func InitLoggerConfig(env *library.Env) {
	LoggerConfigs = []*library.LoggerConfig{
		{
			Receiver: &providers.DefaultLogger,
			Name:           "default-log",
			Level:          env.GetString("LOG_DEFAULT_LEVEL"),          //日志级别
			Path:           env.GetString("LOG_DEFAULT_PATH"),           //日志路径
			MaxAgeDay:      env.GetInt("LOG_DEFAULT_MAX_AGE_DAY"),       //最大保存天数
			MaxFileSize:    env.GetInt("LOG_DEFAULT_MAX_FILE_SIZE"),     //单日志文件最大大小M
			MaxBackups:     env.GetInt("LOG_DEFAULT_MAX_BACKUPS"),       //最大历史文件个数
			DebugMode:      env.GetBool("LOG_DEFAULT_DEBUG_MODE"),       //是否开启调试模式（会记录日志点文件和行号）
			PrintInConsole: env.GetBool("LOG_DEFAULT_PRINT_IN_CONSOLE"), //在控制台显示
			CompressFile:   true,                                        //压缩文件
			ServiceName:    env.GetString("SERVICE_NAME"),               //系统服务名
		},
		{
			Receiver: &providers.RequestLogger,
			Name:           "request-log",                               //日志名称
			Level:          env.GetString("LOG_REQUEST_LEVEL"),          //日志级别
			Path:           env.GetString("LOG_REQUEST_PATH"),           //日志路径
			MaxAgeDay:      env.GetInt("LOG_REQUEST_MAX_AGE_DAY"),       //最大保存天数
			MaxFileSize:    env.GetInt("LOG_REQUEST_MAX_FILE_SIZE"),     //单日志文件最大大小M
			MaxBackups:     env.GetInt("LOG_REQUEST_MAX_BACKUPS"),       //最大历史文件个数
			DebugMode:      env.GetBool("LOG_REQUEST_DEBUG_MODE"),       //是否开启调试模式（会记录日志点文件和行号）
			PrintInConsole: env.GetBool("LOG_REQUEST_PRINT_IN_CONSOLE"), //在控制台显示
			CompressFile:   true,                                        //压缩文件
			ServiceName:    env.GetString("SERVICE_NAME"),               //系统服务名
		},
		{
			Receiver: &providers.CallApiLogger,
			Name:           "call-api-log",                               //日志名称
			Level:          env.GetString("LOG_API_LEVEL"),          //日志级别
			Path:           env.GetString("LOG_API_PATH"),           //日志路径
			MaxAgeDay:      env.GetInt("LOG_API_MAX_AGE_DAY"),       //最大保存天数
			MaxFileSize:    env.GetInt("LOG_API_MAX_FILE_SIZE"),     //单日志文件最大大小M
			MaxBackups:     env.GetInt("LOG_API_MAX_BACKUPS"),       //最大历史文件个数
			DebugMode:      env.GetBool("LOG_API_DEBUG_MODE"),       //是否开启调试模式（会记录日志点文件和行号）
			PrintInConsole: env.GetBool("LOG_API_PRINT_IN_CONSOLE"), //在控制台显示
			CompressFile:   true,                                        //压缩文件
			ServiceName:    env.GetString("SERVICE_NAME"),               //系统服务名
		},
	}
}
