package biz

import (
	cpuv1 "github.com/pachirode/monitor/internal/apiserver/biz/v1/cpu"
	diskv1 "github.com/pachirode/monitor/internal/apiserver/biz/v1/disk"
	memoryv1 "github.com/pachirode/monitor/internal/apiserver/biz/v1/memory"
	networkv1 "github.com/pachirode/monitor/internal/apiserver/biz/v1/network"
	"github.com/pachirode/monitor/internal/apiserver/pkg/monitors"
)

type IBiz interface {
	CpuV1() cpuv1.CpuBiz
	DiskV1() diskv1.DiskBiz
	MemoryV1() memoryv1.MemoryBiz
	NetworkV1() networkv1.NetworkBiz
}

type biz struct {
	manager *monitors.MonitorManager
}

var _ IBiz = (*biz)(nil)

func NewBiz() *biz {
	return &biz{manager: monitors.NewMonitorManager()}
}

func (b *biz) CpuV1() cpuv1.CpuBiz {
	return cpuv1.New(b.manager)
}

func (b *biz) DiskV1() diskv1.DiskBiz {
	return diskv1.New(b.manager)
}

func (b *biz) MemoryV1() memoryv1.MemoryBiz {
	return memoryv1.New(b.manager)
}

func (b *biz) NetworkV1() networkv1.NetworkBiz {
	return networkv1.New(b.manager)
}
