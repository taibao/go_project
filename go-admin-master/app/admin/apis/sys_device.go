package apis

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	"github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth/user"
	_ "github.com/go-admin-team/go-admin-core/sdk/pkg/response"

	"go-admin/app/admin/models"
	"go-admin/app/admin/service"
	"go-admin/app/admin/service/dto"
	"go-admin/common/actions"
)

type SysDevice struct {
	api.Api
}

// GetDeviceList 获取设备列表
func (e SysDevice) GetDeviceList(c *gin.Context) {
	err := e.MakeContext(c).
		Errors
	if err != nil {
		e.Logger.Error(err)
		return
	}
	e.OK("hello world ！", "success")
}

// GetPage 获取SysDevice列表
// @Summary 获取SysDevice列表
// @Description 获取SysDevice列表
// @Tags SysDevice
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.SysDevice}} "{"code": 200, "data": [...]}"
// @Router /api/v1/sys-device [get]
// @Security Bearer
func (e SysDevice) GetPage(c *gin.Context) {
	req := dto.SysDeviceGetPageReq{}
	s := service.SysDevice{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	p := actions.GetPermissionFromContext(c)
	list := make([]models.SysDevice, 0)
	var count int64

	fmt.Printf("输出内容", req)

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取SysDevice失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取SysDevice
// @Summary 获取SysDevice
// @Description 获取SysDevice
// @Tags SysDevice
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.SysDevice} "{"code": 200, "data": [...]}"
// @Router /api/v1/sys-device/{id} [get]
// @Security Bearer
func (e SysDevice) Get(c *gin.Context) {
	req := dto.SysDeviceGetReq{}
	s := service.SysDevice{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	var object models.SysDevice

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取SysDevice失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(object, "查询成功")
}

// Insert 创建SysDevice
// @Summary 创建SysDevice
// @Description 创建SysDevice
// @Tags SysDevice
// @Accept application/json
// @Product application/json
// @Param data body dto.SysDeviceInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/sys-device [post]
// @Security Bearer
func (e SysDevice) Insert(c *gin.Context) {
	req := dto.SysDeviceInsertReq{}
	s := service.SysDevice{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	// 设置创建人
	req.SetCreateBy(user.GetUserId(c))

	err = s.Insert(&req)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("创建SysDevice失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改SysDevice
// @Summary 修改SysDevice
// @Description 修改SysDevice
// @Tags SysDevice
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.SysDeviceUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/sys-device/{id} [put]
// @Security Bearer
func (e SysDevice) Update(c *gin.Context) {
	req := dto.SysDeviceUpdateReq{}
	s := service.SysDevice{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	req.SetUpdateBy(user.GetUserId(c))
	p := actions.GetPermissionFromContext(c)

	err = s.Update(&req, p)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("修改SysDevice失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "修改成功")
}

// Delete 删除SysDevice
// @Summary 删除SysDevice
// @Description 删除SysDevice
// @Tags SysDevice
// @Param data body dto.SysDeviceDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/sys-device [delete]
// @Security Bearer
func (e SysDevice) Delete(c *gin.Context) {
	s := service.SysDevice{}
	req := dto.SysDeviceDeleteReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	// req.SetUpdateBy(user.GetUserId(c))
	p := actions.GetPermissionFromContext(c)

	err = s.Remove(&req, p)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("删除SysDevice失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "删除成功")
}
