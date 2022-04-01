package repository

import (
	"context"
	"app_es_service_go/data/model"
	"app_es_service_go/providers"
	"fmt"
	"go.uber.org/zap"
)

func GetAppConf(ctx context.Context, appId string) *model.AppConf {
	//rst := map[string]interface{}{}
	r, e := providers.HttpClient.Get(ctx, "https://api.weixin.qq.com/cgi-bin/clear_quota?", nil, providers.CallApiLogger)
	fmt.Println("get function", r, e)
	r2, e2 := providers.HttpClient.Post(ctx, "https://api.weixin.qq.com/cgi-bin/clear_quota?", nil, nil, providers.CallApiLogger)
	fmt.Println("post function", r2, e2)
	appConf := &model.AppConf{}
	providers.CoreGorm.Where("app_id", appId).First(appConf)
	providers.DefaultLogger.Info("query from db", zap.Any("rst", appConf))
	return appConf
}

//调用接口查询店铺列表
func GetShopListFromEsApi(ctx context.Context, appId string) *model.AppConf {
	//r2, e2 := providers.HttpClient.Post(ctx, "https://api.weixin.qq.com/cgi-bin/clear_quota?", nil, nil, providers.CallApiLogger)
	//fmt.Println("post function", r2, e2)
	//appConf := &model.AppConf{}
	//providers.CoreGorm.Where("app_id", appId).First(appConf)
	//providers.DefaultLogger.Info("query from db", zap.Any("rst", appConf))
	//return appConf
}
