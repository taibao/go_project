package library


type GormConfig struct {
	Receiver **GormDB
	ConnectionName string
	DBName string
	Host string
	Port string
	UserName string
	Password string
	MaxLifeTime int
	MaxOpenConn int
	MaxIdleConn int
}
