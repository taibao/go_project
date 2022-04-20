package controller

import (
	"app_es_service_go/http/controller/request"
	"app_es_service_go/providers"
	"app_es_service_go/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func PostShopInfo(ginCtx *gin.Context) {
	req := &request.GetShopInfoRequest{}
	if err := ginCtx.ShouldBind(req); err != nil {
		ginCtx.JSON(200, gin.H{
			"code": 1,
			"msg":  err.Error(),
		})
		return
	}
	providers.DefaultLogger.Info("Get Params", zap.Any("request", req))
	appInfo := service.GetAppInfoByAppId(ginCtx.Request.Context(), req.AppId)
	ginCtx.JSON(200, gin.H{
		"code": 0,
		"msg":  "OK",
		"data": appInfo,
	})
}

func GetShopInfo(ginCtx *gin.Context) {
	AppId := ginCtx.Query("app_id")
	if AppId == ""{
		ginCtx.JSON(200, gin.H{
			"code": 1,
			"msg": "参数错误或缺失",
		})
		return
	}

	providers.DefaultLogger.Info("Get Params", zap.Any("request", "app_id="+AppId))
	appInfo := service.GetAppInfoByAppId(ginCtx.Request.Context(), AppId)
	ginCtx.JSON(200, gin.H{
		"code": 0,
		"msg":  "OK",
		"data": appInfo,
	})
}

/*
*获取全部店铺列表
*/
func GetShopList(ginCtx *gin.Context) {
	ginCtx.JSON(200, gin.H{
		"code": 0,
		"msg":  "OK",
		"data": "",
	})
}
