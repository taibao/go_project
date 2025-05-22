package models

import (
	"go-admin/common/models"
)

type SysDevice struct {
	ID         int64  `gorm:"primaryKey;autoIncrement;comment:设备编号" json:"id"`
	DeviceName string `json:"deviceName" gorm:"size:64;comment:设备名称"`
	UserID     string `json:"userId" gorm:"size:128;comment:用户编号"`
	Status     string `json:"status" gorm:"size:4;comment:状态 0:正常 1：停用 2：删除"`
	models.ControlBy
	models.ModelTime
}

func (*SysDevice) TableName() string {
	return "sys_device"
}

func (e *SysDevice) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *SysDevice) GetId() interface{} {
	return e.ID
}
