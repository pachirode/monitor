package cpu

import (
	"context"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/pachirode/monitor/internal/apiserver/pkg/conversion"
	"github.com/pachirode/monitor/internal/apiserver/pkg/monitors"
	apiv1 "github.com/pachirode/monitor/pkg/api/apiserver/v1"
)

type CPUBiz interface {
	GetCPUInfo(ctx context.Context, rq *emptypb.Empty) (*apiv1.GetCPUResponse, error)

	CPUExpansion
}

type CPUExpansion interface {
}

type cpuBiz struct {
	manager *monitors.MonitorManager
}

var _ CPUBiz = (*cpuBiz)(nil)

func New() *cpuBiz {
	return &cpuBiz{manager: monitors.NewMonitorManager()}
}

func (b *cpuBiz) GetCPUInfo(ctx context.Context, rq *emptypb.Empty) (*apiv1.GetCPUResponse, error) {
	cpuInfo := b.manager.CpuMonitor.GetCPUInfo()
	return &apiv1.GetCPUResponse{Cpu: conversion.CPUInfoToCPUV1(cpuInfo)}, nil
}
