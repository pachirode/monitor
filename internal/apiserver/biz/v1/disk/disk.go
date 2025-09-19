package disk

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/pachirode/monitor/internal/apiserver/pkg/conversion"
	"github.com/pachirode/monitor/internal/apiserver/pkg/monitors"
	apiv1 "github.com/pachirode/monitor/pkg/api/apiserver/v1"
)

type DiskBiz interface {
	GetDiskInfos(ctx context.Context, rq *emptypb.Empty) (*apiv1.GetDisksResponse, error)

	DiskExpansion
}

type DiskExpansion interface {
}

type diskBiz struct {
	manager *monitors.MonitorManager
}

var _ DiskBiz = (*diskBiz)(nil)

func New(manager *monitors.MonitorManager) *diskBiz {
	return &diskBiz{manager: manager}
}

func (b *diskBiz) GetDiskInfos(ctx context.Context, rq *emptypb.Empty) (*apiv1.GetDisksResponse, error) {
	diskInfos := b.manager.DiskMonitor.GetDiskInfos()
	return &apiv1.GetDisksResponse{DiskList: conversion.DiskInfoToDiskV1(diskInfos)}, nil
}
