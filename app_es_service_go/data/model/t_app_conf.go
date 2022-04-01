package model

import "database/sql"

type AppConf struct {
	AppId string `gorm:"column:app_id"`
	WxAppType int `gorm:"column:wx_app_type"`
	MerchantId string `gorm:"column:merchant_id"`
	WxAppId sql.NullString `gorm:"column:wx_app_id"`
	WxAppName string `gorm:"column:wx_app_name"`
	ShopName string `gorm:"column:shop_name"`
}

func (AppConf) TableName() string {
	return "t_app_conf"
}