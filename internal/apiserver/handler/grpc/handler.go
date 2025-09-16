package grpc

import (
	"github.com/pachirode/monitor/internal/apiserver/biz"
	apiv1 "github.com/pachirode/monitor/pkg/api/apiserver/v1"
)

type Handler struct {
	apiv1.UnimplementedMonitorServer

	biz biz.IBiz
}

func NewHandler(biz biz.IBiz) *Handler {
	return &Handler{
		biz: biz,
	}
}
