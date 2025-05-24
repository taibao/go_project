package router

import (
	"github.com/gin-gonic/gin"
	jwt "github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth"
	"go-admin/app/admin/apis"
	"go-admin/common/actions"
	"go-admin/common/middleware"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerSysDeviceRouter)
}

// 需认证的路由代码
func registerSysDeviceRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	api := apis.SysDevice{}
	r := v1.Group("/sys-Device").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole()).Use(actions.PermissionAction())
	{
		r.GET("", api.GetPage)
		r.GET("/:id", api.Get) //获取设备详情
		r.POST("", api.Insert)
		r.PUT("", api.Update)
		r.DELETE("", api.Delete)
	}

	//Device := v1.Group("/Device").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole()).Use(actions.PermissionAction())
	//{
	//	Device.GET("/profile", api.GetProfile)
	//	Device.POST("/avatar", api.InsetAvatar)
	//	Device.PUT("/pwd/set", api.UpdatePwd)
	//	Device.PUT("/pwd/reset", api.ResetPwd)
	//	Device.PUT("/status", api.UpdateStatus)
	//}

	//设备管理列表
	devices := v1.Group("/devices").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole()).Use(actions.PermissionAction())
	{
		devices.GET("", api.GetPage)
		devices.GET("/:id", api.Get)
		devices.POST("", api.Insert)
		devices.PUT("", api.Update)
		devices.DELETE("", api.Delete)
	}

}
