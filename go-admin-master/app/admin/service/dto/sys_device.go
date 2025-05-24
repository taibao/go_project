package dto

import (
	"go-admin/app/admin/models"

	"go-admin/common/dto"
	common "go-admin/common/models"
)

type SysDeviceGetPageReq struct {
	dto.Pagination `search:"-"`
	Id             int    `form:"id" search:"type:exact;column:id;table:sys_device" comment:"用户id"`
	Username       string `form:"username" search:"type:contains;column:username;table:sys_device" comment:"用户名"`
	UserId         string `form:"user_id" search:"type:contains;column:user_id;table:sys_device" comment:"用户名"`
	DeviceName     string `form:"deviceName" search:"type:contains;column:device_name;table:sys_device" comment:"昵称"`
	Status         string `form:"status" search:"type:exact;column:status;table:sys_device" comment:"状态"`
	CreatedAt      string `form:"created_at" search:"type:exact;column:created_at;table:sys_device" comment:"创建时间"`
}

func (m *SysDeviceGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type SysDeviceById struct {
	dto.ObjectById
	common.ControlBy
}

func (s *SysDeviceById) GetId() interface{} {
	if len(s.Ids) > 0 {
		s.Ids = append(s.Ids, s.Id)
		return s.Ids
	}
	return s.Id
}

func (s *SysDeviceById) GenerateM() (common.ActiveRecord, error) {
	return &models.SysDevice{}, nil
}

type SysDeviceInsertReq struct {
	Id         string `json:"id" gorm:"size:64;comment:设备名称"`
	DeviceName string `json:"device_name" gorm:"size:64;comment:设备名称"`
	UserID     int    `json:"user_id" gorm:"size:128;comment:用户编号"`
	DeviceSn   string `json:"device_sn" gorm:"size:128;comment:设备sn号"`
	ImgUrl     string `json:"img_url" gorm:"size:128;comment:设备图片"`
	Status     string `json:"status" gorm:"size:4;comment:状态 0:正常 1：停用 2：删除"`
	common.ControlBy
}

func (s *SysDeviceInsertReq) GetId() interface{} {
	return s.Id
}

func (s *SysDeviceInsertReq) Generate(model *models.SysDevice) {
	if s.UserID != 0 {
		model.UserID = s.UserID
	}
	model.Status = s.Status
	model.DeviceSn = s.DeviceSn
	model.DeviceName = s.DeviceName
	model.ImgUrl = s.ImgUrl
}

type SysDeviceUpdateReq struct {
	Id         int64  `json:"id" gorm:"size:64;comment:设备名称"`
	DeviceName string `json:"device_name" gorm:"size:64;comment:设备名称"`
	UserID     int    `json:"user_id" gorm:"size:128;comment:用户编号"`
	DeviceSn   string `json:"device_sn" gorm:"size:128;comment:设备sn号"`
	ImgUrl     string `json:"img_url" gorm:"size:128;comment:设备图片"`
	Status     string `json:"status" gorm:"size:4;comment:状态 0:正常 1：停用 2：删除"`
	common.ControlBy
}

func (s *SysDeviceUpdateReq) GetId() interface{} {
	return s.Id
}

func (s *SysDeviceUpdateReq) Generate(model *models.SysDevice) {
	if s.Id != 0 {
		model.ID = s.Id
	}
	model.DeviceName = s.DeviceName
}

type UpdateSysDeviceImageUrlReq struct {
	Id         int64  `json:"id" gorm:"size:64;comment:设备名称"`
	DeviceName string `json:"device_name" gorm:"size:64;comment:设备名称"`
	UserID     int    `json:"user_id" gorm:"size:128;comment:用户编号"`
	DeviceSn   string `json:"device_sn" gorm:"size:128;comment:设备sn号"`
	ImgUrl     string `json:"img_url" gorm:"size:128;comment:设备图片"`
	Status     string `json:"status" gorm:"size:4;comment:状态 0:正常 1：停用 2：删除"`
	common.ControlBy
}

func (s *UpdateSysDeviceImageUrlReq) GetId() interface{} {
	return s.Id
}

// 更新设备状态
type UpdateSysDeviceStatusReq struct {
	Id         int64  `json:"id" gorm:"size:64;comment:设备名称"`
	DeviceName string `json:"device_name" gorm:"size:64;comment:设备名称"`
	UserID     int    `json:"user_id" gorm:"size:128;comment:用户编号"`
	DeviceSn   string `json:"device_sn" gorm:"size:128;comment:设备sn号"`
	ImgUrl     string `json:"img_url" gorm:"size:128;comment:设备图片"`
	Status     string `json:"status" gorm:"size:4;comment:状态 0:正常 1：停用 2：删除"`
	common.ControlBy
}

func (s *UpdateSysDeviceStatusReq) GetId() interface{} {
	return s.Id
}
