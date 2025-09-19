package http

import (
	"github.com/gin-gonic/gin"
	"github.com/pachirode/pkg/core"
	"github.com/pachirode/pkg/log"
)

func (h *Handler) GetCpuInfo(ctx *gin.Context) {
	log.W(ctx.Request.Context()).Infow("GetCpuInfo handler is called", "method", "GetCpuInfo", "status", "Monitor")
	core.HandleJSONRequest(ctx, h.biz.CpuV1().GetCpuInfo)
}
