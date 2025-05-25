package dto

import (

	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type SysDeviceGetPageReq struct {
	dto.Pagination     `search:"-"`
    SysDeviceOrder
}

type SysDeviceOrder struct {
    Id string `form:"idOrder"  search:"type:order;column:id;table:sys_device"`
    DeviceName string `form:"deviceNameOrder"  search:"type:order;column:device_name;table:sys_device"`
    DeviceSn string `form:"deviceSnOrder"  search:"type:order;column:device_sn;table:sys_device"`
    ImgUrl string `form:"imgUrlOrder"  search:"type:order;column:img_url;table:sys_device"`
    UserId string `form:"userIdOrder"  search:"type:order;column:user_id;table:sys_device"`
    Status string `form:"statusOrder"  search:"type:order;column:status;table:sys_device"`
    CreateBy string `form:"createByOrder"  search:"type:order;column:create_by;table:sys_device"`
    UpdateBy string `form:"updateByOrder"  search:"type:order;column:update_by;table:sys_device"`
    CreatedAt string `form:"createdAtOrder"  search:"type:order;column:created_at;table:sys_device"`
    UpdatedAt string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:sys_device"`
    DeletedAt string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:sys_device"`
    
}

func (m *SysDeviceGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type SysDeviceInsertReq struct {
    Id int `json:"-" comment:"设备编号"` // 设备编号
    DeviceName string `json:"deviceName" comment:"设备名称"`
    DeviceSn string `json:"deviceSn" comment:"设备sn号"`
    ImgUrl string `json:"imgUrl" comment:"图片链接"`
    UserId string `json:"userId" comment:"用户编号"`
    Status string `json:"status" comment:"状态 0:正常 1：停用 2：删除"`
    common.ControlBy
}

func (s *SysDeviceInsertReq) Generate(model *models.SysDevice)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.DeviceName = s.DeviceName
    model.DeviceSn = s.DeviceSn
    model.ImgUrl = s.ImgUrl
    model.UserId = s.UserId
    model.Status = s.Status
    model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
}

func (s *SysDeviceInsertReq) GetId() interface{} {
	return s.Id
}

type SysDeviceUpdateReq struct {
    Id int `uri:"id" comment:"设备编号"` // 设备编号
    DeviceName string `json:"deviceName" comment:"设备名称"`
    DeviceSn string `json:"deviceSn" comment:"设备sn号"`
    ImgUrl string `json:"imgUrl" comment:"图片链接"`
    UserId string `json:"userId" comment:"用户编号"`
    Status string `json:"status" comment:"状态 0:正常 1：停用 2：删除"`
    common.ControlBy
}

func (s *SysDeviceUpdateReq) Generate(model *models.SysDevice)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.DeviceName = s.DeviceName
    model.DeviceSn = s.DeviceSn
    model.ImgUrl = s.ImgUrl
    model.UserId = s.UserId
    model.Status = s.Status
    model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
}

func (s *SysDeviceUpdateReq) GetId() interface{} {
	return s.Id
}

// SysDeviceGetReq 功能获取请求参数
type SysDeviceGetReq struct {
     Id int `uri:"id"`
}
func (s *SysDeviceGetReq) GetId() interface{} {
	return s.Id
}

// SysDeviceDeleteReq 功能删除请求参数
type SysDeviceDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *SysDeviceDeleteReq) GetId() interface{} {
	return s.Ids
}
