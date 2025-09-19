package grpc

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	apiv1 "github.com/pachirode/monitor/pkg/api/apiserver/v1"
)

func (h *Handler) GetNetworkInfos(ctx context.Context, rq *emptypb.Empty) (*apiv1.GetNetworksResponse, error) {
	return h.biz.NetworkV1().GetNetworkInfos(ctx, rq)
}
