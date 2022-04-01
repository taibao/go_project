package gin_plugin

import (
	"context"
	"github.com/gin-gonic/gin"
	"talkcheap.xiaoeknow.com/xiaoetong/eframe/contract"
)

func XeSpecificContextSet(ginCtx *gin.Context) {
	ctx := ginCtx.Request.Context()
	xeSpecific := make(map[string]string, 3)
	xeSpecific[contract.XeTagHeader] = ginCtx.GetHeader(contract.XeTagHeader)
	xeSpecific[contract.Sw8Header] = ginCtx.GetHeader(contract.Sw8Header)
	xeSpecific[contract.Sw8CorrelationHeader] = ginCtx.GetHeader(contract.Sw8CorrelationHeader)
	xeCtx := context.WithValue(ctx, contract.XeCtx, xeSpecific)
	ginCtx.Request = ginCtx.Request.WithContext(xeCtx)
}
