package http

import (
	"github.com/gin-gonic/gin"
	"github.com/pachirode/pkg/core"
	"github.com/pachirode/pkg/log"
)

func (h *Handler) GetMemoryInfo(ctx *gin.Context) {
	log.W(ctx.Request.Context()).Infow("GetMemoryInfo handler is called", "method", "GetMemoryInfo", "status", "Monitor")
	core.HandleJSONRequest(ctx, h.biz.MemoryV1().GetMemoryInfo)
}
