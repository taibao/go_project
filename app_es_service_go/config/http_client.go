package config

import "talkcheap.xiaoeknow.com/xiaoetong/eframe/library"

var DefaultHttpClientConfig *library.HttpClientConfig

func InitHttpClientConfig(env *library.Env) {
	DefaultHttpClientConfig = &library.HttpClientConfig{
		Name:                      "http-client-default",
		DialTimeoutSecond:         env.GetInt("HTTP_CLIENT_DIAL_TIMEOUT_SECOND"),
		DialKeepAliveSecond:       env.GetInt("HTTP_CLIENT_DIAL_KEEP_ALIVE_SECOND"),
		MaxIdleConnections:        env.GetInt("HTTP_CLIENT_MAX_IDEL_CONNS"),
		MaxIdleConnectionsPerHost: env.GetInt("HTTP_CLIENT_MAX_IDEL_PER_HOST"),
		IdleConnTimeoutSecond:     env.GetInt("HTTP_CLIENT_MAX_CONN_TIMEOUT_SECOND"),
	}
}
