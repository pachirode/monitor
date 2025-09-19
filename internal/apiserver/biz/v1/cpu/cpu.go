package cpu

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/pachirode/monitor/internal/apiserver/pkg/conversion"
	"github.com/pachirode/monitor/internal/apiserver/pkg/monitors"
	apiv1 "github.com/pachirode/monitor/pkg/api/apiserver/v1"
)

type CpuBiz interface {
	GetCpuInfo(ctx context.Context, rq *emptypb.Empty) (*apiv1.GetCPUResponse, error)

	CpuExpansion
}

type CpuExpansion interface {
}

type cpuBiz struct {
	manager *monitors.MonitorManager
}

var _ CpuBiz = (*cpuBiz)(nil)

func New(manager *monitors.MonitorManager) *cpuBiz {
	return &cpuBiz{manager: manager}
}

func (b *cpuBiz) GetCpuInfo(ctx context.Context, rq *emptypb.Empty) (*apiv1.GetCPUResponse, error) {
	cpuInfo := b.manager.CpuMonitor.GetCpuInfo()
	return &apiv1.GetCPUResponse{Cpu: conversion.CPUInfoToCPUV1(cpuInfo)}, nil
}
