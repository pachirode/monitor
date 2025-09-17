package http

import (
	"github.com/gin-gonic/gin"
	"github.com/pachirode/pkg/core"
	"github.com/pachirode/pkg/log"
)

func (h *Handler) CPUMonitor(c *gin.Context) {
	log.W(c.Request.Context()).Infow("CPUMonitor handler is called", "method", "CPUMonitor", "status", "Monitor")
	core.HandleJSONRequest(c, h.manager)
}
