package service

import (
	"errors"

    "github.com/go-admin-team/go-admin-core/sdk/service"
	"gorm.io/gorm"

	"go-admin/app/admin/models"
	"go-admin/app/admin/service/dto"
	"go-admin/common/actions"
	cDto "go-admin/common/dto"
)

type SysDevice struct {
	service.Service
}

// GetPage 获取SysDevice列表
func (e *SysDevice) GetPage(c *dto.SysDeviceGetPageReq, p *actions.DataPermission, list *[]models.SysDevice, count *int64) error {
	var err error
	var data models.SysDevice

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("SysDeviceService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取SysDevice对象
func (e *SysDevice) Get(d *dto.SysDeviceGetReq, p *actions.DataPermission, model *models.SysDevice) error {
	var data models.SysDevice

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetSysDevice error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建SysDevice对象
func (e *SysDevice) Insert(c *dto.SysDeviceInsertReq) error {
    var err error
    var data models.SysDevice
    c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("SysDeviceService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改SysDevice对象
func (e *SysDevice) Update(c *dto.SysDeviceUpdateReq, p *actions.DataPermission) error {
    var err error
    var data = models.SysDevice{}
    e.Orm.Scopes(
            actions.Permission(data.TableName(), p),
        ).First(&data, c.GetId())
    c.Generate(&data)

    db := e.Orm.Save(&data)
    if err = db.Error; err != nil {
        e.Log.Errorf("SysDeviceService Save error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权更新该数据")
    }
    return nil
}

// Remove 删除SysDevice
func (e *SysDevice) Remove(d *dto.SysDeviceDeleteReq, p *actions.DataPermission) error {
	var data models.SysDevice

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
        e.Log.Errorf("Service RemoveSysDevice error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权删除该数据")
    }
	return nil
}
