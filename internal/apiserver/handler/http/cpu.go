package http

import (
	"github.com/gin-gonic/gin"
	"github.com/pachirode/pkg/core"
	"github.com/pachirode/pkg/log"
)

func (h *Handler) GetCPUInfo(ctx *gin.Context) {
	log.W(ctx.Request.Context()).Infow("GetCPUInfo handler is called", "method", "GetCPUInfo", "status", "Monitor")
	core.HandleJSONRequest(ctx, h.biz.CpuV1().GetCPUInfo)
}
