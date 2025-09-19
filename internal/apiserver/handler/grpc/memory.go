package grpc

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	apiv1 "github.com/pachirode/monitor/pkg/api/apiserver/v1"
)

func (h *Handler) GetMemoryInfo(ctx context.Context, rq *emptypb.Empty) (*apiv1.GetMemoryResponse, error) {
	return h.biz.MemoryV1().GetMemoryInfo(ctx, rq)
}
