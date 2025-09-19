package network

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/pachirode/monitor/internal/apiserver/pkg/conversion"
	"github.com/pachirode/monitor/internal/apiserver/pkg/monitors"
	apiv1 "github.com/pachirode/monitor/pkg/api/apiserver/v1"
)

type NetworkBiz interface {
	GetNetworkInfos(ctx context.Context, rq *emptypb.Empty) (*apiv1.GetNetworksResponse, error)

	NetworkExpansion
}

type NetworkExpansion interface {
}

type networkBiz struct {
	manager *monitors.MonitorManager
}

var _ NetworkBiz = (*networkBiz)(nil)

func New(manager *monitors.MonitorManager) *networkBiz {
	return &networkBiz{manager: manager}
}

func (b *networkBiz) GetNetworkInfos(ctx context.Context, rq *emptypb.Empty) (*apiv1.GetNetworksResponse, error) {
	networkInfos := b.manager.NetworkMonitor.GetNetworkInfos()
	return &apiv1.GetNetworksResponse{NetworkList: conversion.NetworkInfoToNetworkV1(networkInfos)}, nil
}
