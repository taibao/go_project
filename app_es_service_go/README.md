**app_es_service_go**

购

小鹅Go框架使用demo

编译 `go build -o edemo ./cmd`

运行 `edemo(edemo.exe) run -c .env.local -p 1234 -s shop_msg_consume`

`-c` 指定env文件
`-p` 指定http服务运行端口
`-s` 指定需要运行的脚本 多个脚本用`,`分隔 如`shop_msg_consume,material_msg_consume`


```
如果遇到以下报错
fatal: could not read Username for 'https://talkcheap.xiaoeknow.com': terminal prompts disabled
将gitlab代码仓库设置为私有仓库
go env -w GOPRIVATE=talkcheap.xiaoeknow.com
```

### 1.目录结构

```
.
├── README.md
├── bootstrap #启动组件初始化
│   └── app_init.go
├── cmd       #启动入口
│   └── main.go
├── config    #配置项
│   ├── gorm.go   
│   ├── http_client.go
│   ├── http_server.go
│   ├── job.go
│   ├── kafka_group_consumer.go
│   ├── kafka_producer.go
│   ├── log.go
│   ├── mysql.go
│   └── redis.go
├── data    #数据源
│   ├── http_client
│   └── model
│       ├── t_app_conf.go
│       └── t_app_platform_conf.go
├── go.mod
├── go.sum
├── http   #gin路由和控制器
│   ├── controller
│   │   ├── request
│   │   │   └── get_shop_info_request.go
│   │   └── shop_controller.go
│   └── router.go
├── job   #脚本任务
│   └── shop_msg_consume.go
├── providers #服务提供者
│   └── kernel_providers.go
├── repository #数据仓库
│   └── app_conf_repo.go
├── service    #系统服务
│   └── shop_service.go
└── storage   #动态生成文件
    └── logs
        ├── default-log.log
        ├── log-default.log
        ├── log-request.log
        └── request-log.log
```

### 2.核心组件

#### 1.1 Gorm使用

配置说明

```go
type GormConfig struct {
	Receiver **GormDB  //Gorm实例赋值对象
  ConnectionName string //连接名称，可自定义 推荐以gorm打头,如:gorm-core
	DBName string  //数据库名
	Host string    //数据库链接
	Port string		 //数据库端口
  UserName string //用户名称
	Password string //密码
	MaxLifeTime int //设置链接的最大空闲时间。
	MaxOpenConn int //设置最大打开数据库链接数
	MaxIdleConn int //设置最大空闲链接数
}
```

1.在`providers/kernel_providers.go`中，声明对象

```go
import (
	"talkcheap.xiaoeknow.com/xiaoetong/eframe/library"
)

var CoreGorm *library.GormDB
```

2.在`config/gorm.go`中，添加配置


```go
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
	}
}
```

初始化完成后，可在代码逻辑中使用`providers.CoreGorm`调用`gorm.DB`的接口

```go
appConf := &model.AppConf{}
providers.CoreGorm.Where("app_id", appId).First(appConf)
```

#### 1.2 Http-Client使用

1.在`providers/kernel_providers.go`中，声明对象

```go
import (
	"talkcheap.xiaoeknow.com/xiaoetong/eframe/library"
)

var HttpClient *library.HttpClient
```

2.在`config/http_client.go`中，添加配置

```go
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
```

初始化完成后，可在代码逻辑中使用`providers.HttpClient`调用接口



特殊说明

为了skywalking和xe-tag的灰度功能能够正常使用，建议将上一级的`context`透传，因为在`gin`路由初始化的时候，有加入通过`context`进行值传递的逻辑。

```go
import "talkcheap.xiaoeknow.com/xiaoetong/eframe/plugin/gin_plugin"

func LoadRouter(env *library.Env) (router *gin.Engine) {
...
	//容器化灰度标识
	router.Use(gin_plugin.XeSpecificContextSet)
...
}
```

`gin_plugin.XeSpecificContextSet`封装逻辑，将header头中获取的`skywaling`和`xe-tag灰度`标识的值放入到`context`中

```go
package gin_plugin

import (
	"context"
	"github.com/gin-gonic/gin"
	"talkcheap.xiaoeknow.com/xiaoetong/eframe/contract"
)

func XeSpecificContextSet(ginCtx *gin.Context) {
	ctx := ginCtx.Request.Context()
	xeSpecific := make(map[string]string, 3)
	xeSpecific[contract.XeTagHeader] = ginCtx.GetHeader(contract.XeTagHeader)
	xeSpecific[contract.Sw8Header] = ginCtx.GetHeader(contract.Sw8Header)
	xeSpecific[contract.Sw8CorrelationHeader] = ginCtx.GetHeader(contract.Sw8CorrelationHeader)
	xeCtx := context.WithValue(ctx, contract.XeCtx, xeSpecific)
	ginCtx.Request = ginCtx.Request.WithContext(xeCtx)
}
```

`Http-Client`封装逻辑，从`context`中获取到`skywaling`和`xe-tag灰度`标识的值，重新放入到请求的header头中继续向后端服务传递。

```go
package library

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"talkcheap.xiaoeknow.com/xiaoetong/eframe/contract"
	"talkcheap.xiaoeknow.com/xiaoetong/eframe/helpers/network"
	"time"
)

func (c *HttpClient) GetJson(ctx context.Context, url string, response interface{}, logger contract.XiaoeRequestLoggerInterface) error {
	req, err := http.NewRequest("GET", url, nil)
	var responseBytes []byte
	var clientResp *http.Response
	var params string
	var beginTime = time.Now()
	defer recordLog(req, &clientResp, &params, &responseBytes, &err, beginTime, logger)
	if err != nil {
		return err
	}

	addJsonHeader(req)
	addXeHeader(ctx, req)

	clientResp, err = c.Do(req)

	if err != nil {
		err = fmt.Errorf("response is nil: %w", err)
		return err
	}

	if clientResp == nil {
		err = fmt.Errorf("response is nil: %w", err)
		return err
	}
	defer clientResp.Body.Close()

	if clientResp.StatusCode != http.StatusOK {
		err = &contract.HttpResponseError{
			Code: clientResp.StatusCode,
			Msg:  fmt.Sprintf("response error, code %d", clientResp.StatusCode),
		}
		return err
	}

	responseBytes, err = getBytesFromHttpResponse(clientResp)
	if err != nil {
		return err
	}

	if response == nil {
		response = &(map[string]interface{}{})
	}

	if err = json.Unmarshal(responseBytes, response); err != nil {
		return err
	}

	return err
}
```

#### 1.3 Http-Server使用

1.在`providers/kernel_providers.go`中，声明对象

```go
var HttpClient *library.HttpClient
```

2.服务启动端口配置。配置项只有`port`端口号，可以通过在项目启动时的命令`-p 端口号`来进行配置，或者通过在`config/http_server.go`中进行配置。优先级:命令指定>env配置

#### 1.4 Log使用

Log的底层是通过`zap`的封装来实现的

配置说明

```go
type LoggerConfig struct {
	Receiver **Log //实例化接收对象指针
	Level string //日志级别
	Name string //日志名称
	Path string //日志路径
	MaxAgeDay int //最大保存天数
	MaxFileSize int //单日志文件最大大小M
	MaxBackups int //最大历史文件个数
	DebugMode bool //是否开启调试模式（会记录日志所在位置文件和行号）
	PrintInConsole bool //在控制台显示
	CompressFile bool //压缩文件
	ServiceName string //系统服务名
}
```

1.在`providers/kernel_providers`中声明日志实例的接受对象,如

```go
var CallApiLogger *library.Log
```

2.在`config/log.go`中添加新的配置信息

```go
package config

import (
	"app_es_service_go/providers"
	"talkcheap.xiaoeknow.com/xiaoetong/eframe/library"
)

var LoggerConfigs []*library.LoggerConfig

func InitLoggerConfig(env *library.Env) {
        LoggerConfigs = []*library.LoggerConfig{
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
        }
  	}
}
```

#### 1.5 KafkaProducer使用

配置说明

```go
type KafkaProducerConfig struct {
	Receiver      **KafkaSyncProducer //实例接受对象
	Name          string	//连接名称(自定义)
	Version       string    //版本
	BrokerAddress []string	//消息代理服务器地址
	ConsoleDebug  bool //是否进入命令行终端debug模式
	ExtraConfig   *sarama.Config //额外的配置项
}
```

1.在`providers/kernel_providers.go`中声明对象

```go
var KafkaShopMsgSyncProducer *library.KafkaSyncProducer
```

2.在`config/kafka_producer.go`中添加配置

```go
package config

import (
	"app_es_service_go/providers"
	"strings"
	"talkcheap.xiaoeknow.com/xiaoetong/eframe/library"
)

var KafkaSyncProducerConfig []*library.KafkaProducerConfig

func InitKafkaSyncConfigs(env *library.Env) {
	KafkaSyncProducerConfig = []*library.KafkaProducerConfig{
		{
			Receiver:      &providers.KafkaShopMsgSyncProducer,
			Name:          "kafka-shop-msg-sync-producer",
			Version:       env.GetString("KAFKA_FOR_SHOP_MSG_VERSION"),
			BrokerAddress: strings.Split(env.GetString("KAFKA_FOR_SHOP_MSG_BROKER"), ","),
			ConsoleDebug:  true,
		},
	}
}
```

#### 1.6 KafkaGroupConsumer使用

配置说明

```go
type KafkaGroupConsumerConfig struct {
	Receiver      **KafkaGroupConsumer //实例接受对象
	Name          string               //名称（自定义）
	GroupName     string               //消费组名称
	Version       string               //版本
	Topics        []string             //消费主题
	BrokerAddress []string             //消息代理服务器地址
	ConsoleDebug  bool                 //开启终端debug模式
	ExtraConfig   *sarama.Config       //额外配置项
}
```

1.在`providers/kernel_providers.go`中声明对象

```go
var KafkaShopMsgGroupConsumer *library.KafkaGroupConsumer
```

2.在`config/kafka_group_consumer.go`中添加配置

```go
package config

import (
	"app_es_service_go/providers"
	"strings"
	"talkcheap.xiaoeknow.com/xiaoetong/eframe/library"
)

var KafkaGroupConsumerConfig []*library.KafkaGroupConsumerConfig

func InitKafkaGroupConsumerConfigs(env *library.Env) {
	KafkaGroupConsumerConfig = []*library.KafkaGroupConsumerConfig{
		{
			Receiver:      &providers.KafkaShopMsgGroupConsumer,
			Name:          "kafka-shop-msg-group-consumer",
			GroupName:     env.GetString("KAFKA_FOR_SHOP_MSG_GROUP_NAME"),
			Version:       env.GetString("KAFKA_FOR_SHOP_MSG_VERSION"),
			BrokerAddress: strings.Split(env.GetString("KAFKA_FOR_SHOP_MSG_BROKER"), ","),
			Topics:        strings.Split(env.GetString("KAFKA_FOR_SHOP_MSG_CONSUMER_TOPICS"), ","),
			ConsoleDebug:  false,
		},
	}
}
```

特殊说明：kafka组消费者通常是结合脚本进行使用，需要实现以下2个方法

1.消息处理函数，负责对kafka推送到消费者的消息进行处理，需要注意在该方法中进行逻辑处理时长不能超过`Consumer.Group.Heartbeat.Interval`设置的心跳时长（默认为`3s`）。超过后会因为服务端认为组成员下线而进行消费分区的重新分配。

```go
type MessageHandleFun func(message *sarama.ConsumerMessage)
```

2.异常处理函数

```go
type ConsumeErrHandleFunc func(error)
```

然后再开始消费，`StartConsume(ctx context.Context)`，这里的`context.Context`值需要传入外层的ctx上下文，不然会在服务停止时出现kafka消费脚本未能正常停止、消息消费异常的情况。

如：

```go
import (
	"context"
	"app_es_service_go/providers"
	"fmt"
	"github.com/Shopify/sarama"
	"os"
)

func StartShopMsgConsume(ctx context.Context) {
	providers.KafkaShopMsgGroupConsumer.SetMessageHandleFunc(handleShopInfoMsg)  //设置消息处理函数
	providers.KafkaShopMsgGroupConsumer.SetConsumeErrHandleFunc(handleConsumeError) //设置错误处理函数
	err := providers.KafkaShopMsgGroupConsumer.StartConsume(ctx) //开始消费 
	if err != nil {
		_, _ = fmt.Fprintf(os.Stdout, "%s", err.Error())
	}
}

func handleShopInfoMsg(msg *sarama.ConsumerMessage) {
	fmt.Println("consume msg:", msg.Value)
}

func handleConsumeError(err error) {
	fmt.Println("consume err:", err)
}
```

#### 1.7 Redis使用

配置说明

```go
type RedisConfig struct {
	ConnectionName string //连接名称自定义
	Addr           string //地址
	Port           int    //端口
	Password       string //密码
	DB             int
	PoolSize       int //连接池大小
}
```

1.在`providers/kernel_providers.go`中声明对象

```go
var DefaultRedis *library.RedisClient
```

2.在`config/redis.go`中添加配置信息

```go
package config

import (
	"talkcheap.xiaoeknow.com/xiaoetong/eframe/library"
)

var DefaultRedisConfig *library.RedisConfig

func InitRedisConfig(env *library.Env) {
	DefaultRedisConfig = &library.RedisConfig{
		ConnectionName: "redis-default",
		Addr:     env.GetString("REDIS_DEFAULT_RW_HOST"),
		Port:     env.GetInt("REDIS_DEFAULT_RW_PORT"),
		Password: env.GetString("REDIS_DEFAULT_RW_PASSWORD"),
		DB:       env.GetInt("REDIS_DEFAULT_DATABASE"),
		PoolSize: env.GetInt("REDIS_DEFAULT_POOL_SIZE"),
	}
}
```

#### 1.8 Job使用

目前可以在启动命令中加入`-s` 指定需要运行的脚本 多个脚本用`,`分隔 如`shop_msg_consume,material_msg_consume`

在`job/`目录编写脚本逻辑，然后在`config/job.go`指定脚本名称和对应应的执行函数,如:

```go
package config

import (
	"context"
	"app_es_service_go/job"
)

var JobConfigs = map[string]func(ctx context.Context){
	"shop_msg_consume": job.StartShopMsgConsume,
}

```




