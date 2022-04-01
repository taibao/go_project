package library

type RedisConfig struct {
	ConnectionName string
	Addr string
	Port int
	Password string
	DB int
	PoolSize int
}
