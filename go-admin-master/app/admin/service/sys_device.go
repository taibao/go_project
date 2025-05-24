package service

import (
	"errors"
	"fmt"
	log "github.com/go-admin-team/go-admin-core/logger"
	"github.com/go-admin-team/go-admin-core/sdk/service"
	"go-admin/app/admin/models"
	"go-admin/app/admin/service/dto"
	"go-admin/common/actions"
	cDto "go-admin/common/dto"
	"gorm.io/gorm"
)

type SysDevice struct {
	service.Service
}

// GetPage 获取SysDevice列表
func (e *SysDevice) GetPage(c *dto.SysDeviceGetPageReq, p *actions.DataPermission, list *[]models.SysDevice, count *int64) error {
	var err error
	var data models.SysDevice

	fmt.Println("开始查询----")
	err = e.Orm.Debug().Scopes(
		cDto.MakeCondition(c.GetNeedSearch()),
		cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
		actions.Permission(data.TableName(), p),
	).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("db error: %s", err)
		return err
	}
	return nil
}

// Get 获取SysDevice对象
func (e *SysDevice) Get(d *dto.SysDeviceById, p *actions.DataPermission, model *models.SysDevice) error {
	var data models.SysDevice

	err := e.Orm.Model(&data).Debug().
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("db error: %s", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error: %s", err)
		return err
	}
	return nil
}

// Insert 创建SysDevice对象
func (e *SysDevice) Insert(c *dto.SysDeviceInsertReq) error {
	var err error
	var data models.SysDevice
	var i int64
	err = e.Orm.Model(&data).Where("device_name = ?", c.DeviceName).Count(&i).Error
	if err != nil {
		e.Log.Errorf("db error: %s", err)
		return err
	}
	if i > 0 {
		err := errors.New("设备名称已存在！")
		e.Log.Errorf("db error: %s", err)
		return err
	}
	c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("db error: %s", err)
		return err
	}
	return nil
}

// Update 修改SysDevice对象
func (e *SysDevice) Update(c *dto.SysDeviceUpdateReq, p *actions.DataPermission) error {
	var err error
	var model models.SysDevice
	db := e.Orm.Scopes(
		actions.Permission(model.TableName(), p),
	).First(&model, c.GetId())
	if err = db.Error; err != nil {
		e.Log.Errorf("Service UpdateSysDevice error: %s", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")

	}
	c.Generate(&model)
	update := e.Orm.Model(&model).Where("id = ?", &model.ID).Omit("password", "salt").Updates(&model)
	if err = update.Error; err != nil {
		e.Log.Errorf("db error: %s", err)
		return err
	}
	if update.RowsAffected == 0 {
		err = errors.New("update userinfo error")
		log.Warnf("db update error")
		return err
	}
	return nil
}

// UpdateAvatar 更新设备照
func (e *SysDevice) UpdateAvatar(c *dto.UpdateSysDeviceImageUrlReq, p *actions.DataPermission) error {
	var err error
	var model models.SysDevice
	db := e.Orm.Scopes(
		actions.Permission(model.TableName(), p),
	).First(&model, c.GetId())
	if err = db.Error; err != nil {
		e.Log.Errorf("Service UpdateSysDevice error: %s", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")

	}
	err = e.Orm.Table(model.TableName()).Where("id =? ", c.Id).Updates(c).Error
	if err != nil {
		e.Log.Errorf("Service UpdateSysDevice error: %s", err)
		return err
	}
	return nil
}

// UpdateStatus 更新用户状态
func (e *SysDevice) UpdateStatus(c *dto.UpdateSysDeviceStatusReq, p *actions.DataPermission) error {
	var err error
	var model models.SysDevice
	db := e.Orm.Scopes(
		actions.Permission(model.TableName(), p),
	).First(&model, c.GetId())
	if err = db.Error; err != nil {
		e.Log.Errorf("Service UpdateSysDevice error: %s", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")

	}
	err = e.Orm.Table(model.TableName()).Where("id =? ", c.Id).Updates(c).Error
	if err != nil {
		e.Log.Errorf("Service UpdateSysDevice error: %s", err)
		return err
	}
	return nil
}

// ResetPwd 重置用户密码
//func (e *SysDevice) ResetPwd(c *dto.ResetSysDevicePwdReq, p *actions.DataPermission) error {
//	var err error
//	var model models.SysDevice
//	db := e.Orm.Scopes(
//		actions.Permission(model.TableName(), p),
//	).First(&model, c.GetId())
//	if err = db.Error; err != nil {
//		e.Log.Errorf("At Service ResetSysDevicePwd error: %s", err)
//		return err
//	}
//	if db.RowsAffected == 0 {
//		return errors.New("无权更新该数据")
//	}
//	c.Generate(&model)
//	err = e.Orm.Omit("username", "nick_name", "phone", "role_id", "avatar", "sex").Save(&model).Error
//	if err != nil {
//		e.Log.Errorf("At Service ResetSysDevicePwd error: %s", err)
//		return err
//	}
//	return nil
//}

// Remove 删除SysDevice
func (e *SysDevice) Remove(c *dto.SysDeviceById, p *actions.DataPermission) error {
	var err error
	var data models.SysDevice

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, c.GetId())
	if err = db.Error; err != nil {
		e.Log.Errorf("Error found in  RemoveSysDevice : %s", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权删除该数据")
	}
	return nil
}
