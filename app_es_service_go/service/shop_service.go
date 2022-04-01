package service

import (
	"context"
	"app_es_service_go/repository"
)

type AppInfo struct {
	AppId      string `json:"app_id"`
	MerchantId string `json:"merchant_id"`
	ShopName   string `json:"shop_name"`
}

func GetAppInfoByAppId(ctx context.Context, appId string) *AppInfo {
	appConfModel := repository.GetAppConf(ctx, appId)
	return &AppInfo{
		AppId:      appConfModel.AppId,
		MerchantId: appConfModel.MerchantId,
		ShopName:   appConfModel.ShopName,
	}
}
