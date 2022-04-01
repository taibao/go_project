package library

type LoggerConfig struct {
	Receiver **Log //实例化接收对象指针
	Level string //日志级别
	Name string //日志名称
	Path string //日志路径
	MaxAgeDay int //最大保存天数
	MaxFileSize int //单日志文件最大大小M
	MaxBackups int //最大历史文件个数
	DebugMode bool //是否开启调试模式（会记录日志点文件和行号）
	PrintInConsole bool //在控制台显示
	CompressFile bool //压缩文件
	ServiceName string //系统服务名
}
