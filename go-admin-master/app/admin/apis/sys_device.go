package apis

import (
	"github.com/gin-gonic/gin/binding"
	"go-admin/app/admin/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	"github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth/user"
	_ "github.com/go-admin-team/go-admin-core/sdk/pkg/response"
	"go-admin/app/admin/service"
	"go-admin/app/admin/service/dto"
	"go-admin/common/actions"
)

type SysDevice struct {
	api.Api
}

// GetPage
// @Summary 列表用户信息数据
// @Description 获取JSON
// @Tags 用户
// @Param username query string false "username"
// @Success 200 {string} {object} response.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/sys-device [get]
// @Security Bearer
func (e SysDevice) GetPage(c *gin.Context) {
	s := service.SysDevice{}
	req := dto.SysDeviceGetPageReq{}
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

	//数据权限检查
	p := actions.GetPermissionFromContext(c)

	list := make([]models.SysDevice, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, "查询失败")
		return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get
// @Summary 获取用户
// @Description 获取JSON
// @Tags 用户
// @Param userId path int true "用户编码"
// @Success 200 {object} response.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/sys-device/{userId} [get]
// @Security Bearer
func (e SysDevice) Get(c *gin.Context) {
	s := service.SysDevice{}
	req := dto.SysDeviceById{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req, nil).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	var object models.SysDevice
	//数据权限检查
	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(http.StatusUnprocessableEntity, err, "查询失败")
		return
	}
	e.OK(object, "查询成功")
}

// Insert
// @Summary 创建用户
// @Description 获取JSON
// @Tags 用户
// @Accept  application/json
// @Product application/json
// @Param data body dto.SysDeviceInsertReq true "用户数据"
// @Success 200 {object} response.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/sys-device [post]
// @Security Bearer
func (e SysDevice) Insert(c *gin.Context) {
	s := service.SysDevice{}
	req := dto.SysDeviceInsertReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req, binding.JSON).
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
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update
// @Summary 修改用户数据
// @Description 获取JSON
// @Tags 用户
// @Accept  application/json
// @Product application/json
// @Param data body dto.SysDeviceUpdateReq true "body"
// @Success 200 {object} response.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/sys-device/{userId} [put]
// @Security Bearer

//func (e SysDevice) Update(c *gin.Context) {
//	s := service.SysDevice{}
//	req := dto.SysDeviceUpdateReq{}
//	err := e.MakeContext(c).
//		MakeOrm().
//		Bind(&req).
//		MakeService(&s.Service).
//		Errors
//	if err != nil {
//		e.Logger.Error(err)
//		e.Error(500, err, err.Error())
//		return
//	}
//
//	req.SetUpdateBy(user.GetUserId(c))
//
//	//数据权限检查
//	p := actions.GetPermissionFromContext(c)
//
//	err = s.Update(&req, p)
//	if err != nil {
//		e.Logger.Error(err)
//		return
//	}
//	e.OK(req.GetId(), "更新成功")
//}

// Delete
// @Summary 删除用户数据
// @Description 删除数据
// @Tags 用户
// @Param userId path int true "userId"
// @Success 200 {object} response.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/sys-device/{userId} [delete]
// @Security Bearer

//func (e SysDevice) Delete(c *gin.Context) {
//	s := service.SysDevice{}
//	req := dto.SysDeviceById{}
//	err := e.MakeContext(c).
//		MakeOrm().
//		Bind(&req, binding.JSON).
//		MakeService(&s.Service).
//		Errors
//	if err != nil {
//		e.Logger.Error(err)
//		e.Error(500, err, err.Error())
//		return
//	}
//
//	// 设置编辑人
//	req.SetUpdateBy(user.GetUserId(c))
//
//	// 数据权限检查
//	p := actions.GetPermissionFromContext(c)
//
//	err = s.Remove(&req, p)
//	if err != nil {
//		e.Logger.Error(err)
//		return
//	}
//	e.OK(req.GetId(), "删除成功")
//}
