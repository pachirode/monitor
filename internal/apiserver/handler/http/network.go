package http

import (
	"github.com/gin-gonic/gin"
	"github.com/pachirode/pkg/core"
	"github.com/pachirode/pkg/log"
)

func (h *Handler) GetNetworkInfos(ctx *gin.Context) {
	log.W(ctx.Request.Context()).Infow("GetNetworkInfos handler is called", "method", "GetNetworkInfos", "status", "Monitor")
	core.HandleJSONRequest(ctx, h.biz.NetworkV1().GetNetworkInfos)
}
