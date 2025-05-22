package dto

import (
	"go-admin/app/admin/models"

	"go-admin/common/dto"
	common "go-admin/common/models"
)

type SysDeviceGetPageReq struct {
	dto.Pagination `search:"-"`
	UserId         int    `form:"userId" search:"type:exact;column:user_id;table:sys_device" comment:"用户ID"`
	Username       string `form:"username" search:"type:contains;column:username;table:sys_device" comment:"用户名"`
	NickName       string `form:"nickName" search:"type:contains;column:nick_name;table:sys_device" comment:"昵称"`
	Phone          string `form:"phone" search:"type:contains;column:phone;table:sys_device" comment:"手机号"`
	RoleId         string `form:"roleId" search:"type:exact;column:role_id;table:sys_device" comment:"角色ID"`
	Sex            string `form:"sex" search:"type:exact;column:sex;table:sys_device" comment:"性别"`
	Email          string `form:"email" search:"type:contains;column:email;table:sys_device" comment:"邮箱"`
	PostId         string `form:"postId" search:"type:exact;column:post_id;table:sys_device" comment:"岗位"`
	Status         string `form:"status" search:"type:exact;column:status;table:sys_device" comment:"状态"`
	CreatedAt      string `form:"created_at" search:"type:exact;column:created_at;table:sys_device" comment:"创建时间"`
}

type SysDeviceOrder struct {
	UserIdOrder    string `search:"type:order;column:user_id;table:sys_device" form:"userIdOrder"`
	UsernameOrder  string `search:"type:order;column:username;table:sys_device" form:"usernameOrder"`
	StatusOrder    string `search:"type:order;column:status;table:sys_device" form:"statusOrder"`
	CreatedAtOrder string `search:"type:order;column:created_at;table:sys_device" form:"createdAtOrder"`
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
	UserID     string `json:"user_id" gorm:"size:128;comment:用户编号"`
	DeviceSn   string `json:"device_sn" gorm:"size:128;comment:设备sn号"`
	ImgUrl     string `json:"img_url" gorm:"size:128;comment:设备图片"`
	Status     string `json:"status" gorm:"size:4;comment:状态 0:正常 1：停用 2：删除"`
	common.ControlBy
}

func (s *SysDeviceInsertReq) GetId() interface{} {
	return s.Id
}
