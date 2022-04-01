package library

type MysqlConfig struct {
	Receiver **DB
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