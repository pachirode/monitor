package http

import (
	"github.com/gin-gonic/gin"
	"github.com/pachirode/pkg/core"
	"github.com/pachirode/pkg/log"
)

func (h *Handler) GetDiskInfos(ctx *gin.Context) {
	log.W(ctx.Request.Context()).Infow("GetDiskInfos handler is called", "method", "GetDiskInfos", "status", "Monitor")
	core.HandleJSONRequest(ctx, h.biz.DiskV1().GetDiskInfos)
}
