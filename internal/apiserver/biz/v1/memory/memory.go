package memory

import (
	"context"
	"github.com/pachirode/monitor/internal/apiserver/pkg/conversion"

	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/pachirode/monitor/internal/apiserver/pkg/monitors"
	apiv1 "github.com/pachirode/monitor/pkg/api/apiserver/v1"
)

type MemoryBiz interface {
	GetMemoryInfo(ctx context.Context, rq *emptypb.Empty) (*apiv1.GetMemoryResponse, error)

	MemoryExpansion
}

type MemoryExpansion interface {
}

type memoryBiz struct {
	manager *monitors.MonitorManager
}

var _ MemoryBiz = (*memoryBiz)(nil)

func New(manger *monitors.MonitorManager) *memoryBiz {
	return &memoryBiz{
		manager: manger,
	}
}

func (b *memoryBiz) GetMemoryInfo(ctx context.Context, rq *emptypb.Empty) (*apiv1.GetMemoryResponse, error) {
	memoryInfo := b.manager.GetMemoryInfo()
	return &apiv1.GetMemoryResponse{MemoryInfo: conversion.MemoryInfoToMemoryV1(memoryInfo)}, nil
}
