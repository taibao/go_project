package http

import (
	"app_es_service_go/http/controller"
	"app_es_service_go/providers"
	//"github.com/SkyAPM/go2sky"
	//"github.com/SkyAPM/go2sky-plugins/gin/v3"
	//"github.com/SkyAPM/go2sky/reporter"
	"talkcheap.xiaoeknow.com/xiaoetong/eframe/library"
	"talkcheap.xiaoeknow.com/xiaoetong/eframe/plugin/gin_plugin"
	//"github.com/SkyAPM/go2sky"
	"github.com/gin-gonic/gin"
	"strings"
)

//var tracer *go2sky.Tracer

func LoadRouter(env *library.Env) (router *gin.Engine) {
	gin.SetMode(gin.DebugMode)
	router = gin.New()

	//内网访问中间件初始化
	appEnv := env.GetString("APP_ENV")
	internalIps := strings.Split(",", env.GetString("TENCENT_LB_IP"))
	InternalMiddleware := gin_plugin.InternalMiddleware(appEnv, internalIps)

	//日志中间件初始化
	RequestLogMiddleware := gin_plugin.RequestLogMiddleware(providers.RequestLogger)

	//rp, err := reporter.NewGRPCReporter("addr", reporter.WithCheckInterval(5*time.Second))
	//if err == nil {
	//	tracer, err = go2sky.NewTracer("server_name", go2sky.WithReporter(rp))
	//	if err == nil {
	//		router.Use(v3.Middleware(router, tracer))
	//	}
	//}
	//容器化灰度标识
	router.GET("api/test/gorm", controller.GetShopInfo)
	router.POST("api/test/gorm", controller.PostShopInfo)

	router.Use(gin_plugin.XeSpecificContextSet)
	router.Use(RequestLogMiddleware)
	internalRouter := router.Group("api", InternalMiddleware)
	internalRouter.GET("test/gorm/log", controller.GetShopInfo)

	//获取全部课程列表
	internalRouter.POST("get.shop.list/1.0.0",controller.GetShopList)

	return router
}
