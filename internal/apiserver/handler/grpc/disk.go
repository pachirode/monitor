package grpc

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	apiv1 "github.com/pachirode/monitor/pkg/api/apiserver/v1"
)

func (h *Handler) GetDiskInfos(ctx context.Context, rq *emptypb.Empty) (*apiv1.GetDisksResponse, error) {
	return h.biz.DiskV1().GetDiskInfos(ctx, rq)
}
