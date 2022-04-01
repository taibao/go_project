package bootstrap

import (
	"app_es_service_go/config"
	"app_es_service_go/http"
	"app_es_service_go/providers"
	"context"
	"fmt"
	"github.com/urfave/cli/v2"
	"os"
	"strings"
	"talkcheap.xiaoeknow.com/xiaoetong/eframe/library"
	"talkcheap.xiaoeknow.com/xiaoetong/eframe/library/closable"
)

func Init(ctx context.Context) error {
	cliConf := map[string]interface{}{}
	//cli 启动初始化
	cliApp := &cli.App{
		Commands: []*cli.Command{
			loadRunCommand(&cliConf),
		},
	}
	err := cliApp.Run(os.Args)
	if err != nil {
		return err
	}

	//Env 初始化
	wd, e := os.Getwd()
	if e != nil {
		return err
	}
	if fileName, ok := cliConf["config"]; ok {
		if providers.Env, err = library.NewEnv(fileName.(string), "env", wd); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("请指定运行命令")
	}

	//Log初始化
	config.InitLoggerConfig(providers.Env)
	for _, cfg := range config.LoggerConfigs {
		if cfg.Receiver == nil {
			return fmt.Errorf("[%s] config receiver cannot be nil", cfg.Name)
		}
		*cfg.Receiver = library.NewLogger(cfg)
	}

	//DB 初始化
	config.InitMysqlConfig(providers.Env)
	if providers.CoreDB, err = library.NewDB(config.CoreDBConfig); err != nil {
		return err
	}
	closable.Push(providers.CoreDB)
	if providers.SecondDB, err = library.NewDB(config.ComponentDBConfig); err != nil {
		return err
	}
	closable.Push(providers.SecondDB)

	//Redis 初始化
	config.InitRedisConfig(providers.Env)
	if providers.DefaultRedis, err = library.NewRedisClient(ctx, config.DefaultRedisConfig); err != nil {
		return err
	}
	closable.Push(providers.DefaultRedis)

	if providers.ConfhubRedis, err = library.NewRedisClient(ctx, config.ConfhubRedisConfig); err != nil {
		return err
	}
	closable.Push(providers.ConfhubRedis)

	//GORM 初始化
	config.InitGormConfigs(providers.Env)
	for _, cfg := range config.GormConfigs {
		if cfg.Receiver == nil {
			return fmt.Errorf("[%s] config receiver cannot be nil", cfg.ConnectionName)
		}
		if *cfg.Receiver, err = library.NewGormDB(cfg); err != nil {
			return err
		}
		rawDB, e := (*cfg.Receiver).DB.DB()
		if e != nil {
			return e
		}
		closable.Push(rawDB)
	}

	//Kafka producer 初始化
	//config.InitKafkaSyncConfigs(providers.Env)
	//for _, cfg := range config.KafkaSyncProducerConfig {
	//	if cfg.Receiver == nil {
	//		return fmt.Errorf("[%s] config receiver cannot be nil", cfg.Name)
	//	}
	//	*cfg.Receiver, err = library.NewKafkaSyncProducer(cfg)
	//	if err != nil {
	//		return err
	//	}
	//	closable.PushCloseFun((*cfg.Receiver).Close)
	//}

	//Kafka group consumer 初始化
	//config.InitKafkaGroupConsumerConfigs(providers.Env)
	//for _, cfg := range config.KafkaGroupConsumerConfig {
	//	if cfg.Receiver == nil {
	//		return fmt.Errorf("[%s] config receiver cannot be nil", cfg.Name)
	//	}
	//	*cfg.Receiver, err = library.NewKafkaGroupConsumer(cfg)
	//	if err != nil {
	//		return err
	//	}
	//	closable.PushCloseFun((*cfg.Receiver).Close)
	//}

	//http client 初始化
	config.InitHttpClientConfig(providers.Env)
	providers.HttpClient = library.NewHttpClient(config.DefaultHttpClientConfig)
	closable.PushCloseFun(providers.HttpClient.Close)

	//http server初始化
	providers.Env.Set("HTTP_SERVER_PORT", cliConf["port"].(int))
	config.InitHttpServerConfig(providers.Env)
	server := library.NewHttpServer(config.HttpServerConfig, http.LoadRouter(providers.Env))
	if err = server.Run(); err != nil {
		return err
	}
	closable.Push(server)

	//job 初始化
	scripts := strings.Split(cliConf["script"].(string), ",")
	for _, scriptName := range scripts {
		if f, ok := config.JobConfigs[scriptName]; ok {
			f(ctx)
			_, _ = fmt.Fprintf(os.Stdout, "start script [%s]", scriptName)
		}
	}

	return nil
}

func loadRunCommand(conf *map[string]interface{}) *cli.Command {
	return &cli.Command{
		Name:  "run",
		Usage: "run material center data application",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "config",
				Value:   ".env.local",
				Aliases: []string{"c"},
			},
			&cli.IntFlag{
				Name:    "port",
				Value:   2345,
				Aliases: []string{"p"},
			},
			&cli.StringFlag{
				Name:    "script",
				Value:   "",
				Aliases: []string{"s"},
			},
		},
		Action: loadRunCommandAction(conf),
	}
}

func loadRunCommandAction(conf *map[string]interface{}) func(ctx *cli.Context) error {
	return func(context *cli.Context) error {
		(*conf)["config"] = context.String("c")
		(*conf)["port"] = context.Int("p")
		(*conf)["script"] = context.String("s")
		return nil
	}
}
