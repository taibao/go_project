package models

import (
	"go-admin/common/models"
)

type SysDevice struct {
	models.Model

	DeviceName string `json:"device_name" gorm:"type:varchar(64);comment:设备名称"`
	DeviceSn   string `json:"device_sn" gorm:"type:varchar(64);comment:设备sn号"`
	ImgUrl     string `json:"img_url" gorm:"type:int;comment:图片链接"`
	UserId     string `json:"user_id" gorm:"type:varchar(128);comment:用户编号"`
	Status     string `json:"status" gorm:"type:varchar(4);comment:状态 0:正常 1：停用 2：删除"`
	models.ModelTime
	models.ControlBy
}

func (SysDevice) TableName() string {
	return "sys_device"
}

func (e *SysDevice) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *SysDevice) GetId() interface{} {
	return e.Id
}