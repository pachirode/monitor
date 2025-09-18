package biz

import (
	cpuv1 "github.com/pachirode/monitor/internal/apiserver/biz/v1/cpu"
)

type IBiz interface {
	CpuV1() cpuv1.CPUBiz
}

type biz struct {
}

var _ IBiz = (*biz)(nil)

func NewBiz() *biz {
	return &biz{}
}

func (b *biz) CpuV1() cpuv1.CPUBiz {
	return cpuv1.New()
}
