package library

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"talkcheap.xiaoeknow.com/xiaoetong/eframe/contract"
)

type Log struct {
	*zap.Logger
}

func (log *Log) HttpRequestLog(record *contract.XiaoeHttpRequestRecord) {
	log.Info(record.Msg,
		zap.String("app_id", record.AppId),
		zap.String("sw8", record.Sw8),
		zap.String("sw8_correlation", record.Sw8Correlation),
		zap.String("xe_tag", record.XeTag),
		zap.String("uid", record.Uid),
		zap.Int("http_status", record.HttpStatus),
		zap.String("target_url", record.TargetUrl),
		zap.String("method", record.Method),
		zap.String("client_ip", record.ClientIp),
		zap.String("server_ip", record.ServerIp),
		zap.String("user_agent", record.UserAgent),
		zap.String("begin_time", record.BeginTime),
		zap.String("end_time", record.EndTime),
		zap.Int("cost_time", record.CostTime),
		zap.String("params", record.Params),
		zap.String("response", record.Response))
}

func NewLogger(config *LoggerConfig) *Log {
	encoderConf := zapcore.EncoderConfig{
		MessageKey:       "msg",
		LevelKey:         "level",
		TimeKey:          "log_at",
		NameKey:          config.Name,
		CallerKey:        "caller",
		FunctionKey:      "function",
		StacktraceKey:    "stack",
		SkipLineEnding:   false,
		LineEnding:       zapcore.DefaultLineEnding,
		EncodeLevel:      zapcore.LowercaseLevelEncoder,
		EncodeTime:       zapcore.ISO8601TimeEncoder,
		EncodeDuration:   zapcore.SecondsDurationEncoder,
		EncodeCaller:     zapcore.FullCallerEncoder,
		EncodeName:       zapcore.FullNameEncoder,
		ConsoleSeparator: "",
	}

	atomicLevel := zap.NewAtomicLevel()
	switch config.Level {
	case "info":
		atomicLevel.SetLevel(zap.InfoLevel)
	case "warning":
		atomicLevel.SetLevel(zap.WarnLevel)
	case "error":
		atomicLevel.SetLevel(zap.ErrorLevel)
	default:
		atomicLevel.SetLevel(zap.InfoLevel)
	}

	fileName := fmt.Sprintf("%s/%s.log", config.Path, config.Name)

	hook := lumberjack.Logger{
		Filename:   fileName,
		MaxSize:    config.MaxFileSize,
		MaxAge:     config.MaxAgeDay,
		MaxBackups: config.MaxBackups,
		LocalTime:  true,
		Compress:   config.CompressFile,
	}

	var ws []zapcore.WriteSyncer
	if config.PrintInConsole {
		ws = append(ws, zapcore.AddSync(os.Stdout))
	}
	ws = append(ws, zapcore.AddSync(&hook))

	core := zapcore.NewCore(zapcore.NewJSONEncoder(encoderConf),
		zapcore.NewMultiWriteSyncer(ws...), atomicLevel)

	var options []zap.Option

	if config.DebugMode {
		options = append(options, zap.AddCaller())
		options = append(options, zap.Development())
	}

	filed := zap.Fields(zap.String("service_name", config.ServiceName))

	options = append(options, filed)

	return &Log{
		zap.New(core, options...),
	}
}
