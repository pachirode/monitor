package http

import (
	"github.com/pachirode/monitor/internal/apiserver/biz"
	"github.com/pachirode/monitor/internal/apiserver/pkg/monitors"
	"github.com/pachirode/monitor/internal/apiserver/pkg/validation"
)

type Handler struct {
	biz     biz.IBiz
	val     *validation.Validator
	manager *monitors.MonitorManager
}

func NewHandler(biz biz.IBiz, val *validation.Validator) *Handler {
	return &Handler{
		biz:     biz,
		val:     val,
		manager: monitors.NewMonitorManager(),
	}
}
