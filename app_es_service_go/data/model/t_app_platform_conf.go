package model

type AppPlatformConf struct {
	WxAppId string `gorm:"column:wx_app_id"`
	PlatformType int `gorm:"column:platform_type"`
	CreatedAt string `gorm:"column:created_at"`
	UpdatedAt string `gorm:"column:updated_at"`
}

func (AppPlatformConf) TableName() string  {
	return "t_app_platform_conf"
}
