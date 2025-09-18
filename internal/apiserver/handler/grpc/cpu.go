package grpc

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	apiv1 "github.com/pachirode/monitor/pkg/api/apiserver/v1"
)

func (h *Handler) GetCPUInfo(ctx context.Context, rq *emptypb.Empty) (*apiv1.GetCPUResponse, error) {
	return h.biz.CpuV1().GetCPUInfo(ctx, rq)
}
