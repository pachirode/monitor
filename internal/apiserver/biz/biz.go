package biz

import "github.com/pachirode/monitor/internal/apiserver/pkg/monitors"

type IBiz interface {
	ManagerV1() *monitors.MonitorManager
}

type biz struct {
}

var _ IBiz = (*biz)(nil)

func NewBiz() *biz {
	return &biz{}
}

func (b *biz) ManagerV1() *monitors.MonitorManager {
	return monitors.NewMonitorManager()
}
