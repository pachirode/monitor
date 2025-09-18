package conversion

import (
	"github.com/pachirode/pkg/core"

	"github.com/pachirode/monitor/internal/apiserver/pkg/monitors"
	apiv1 "github.com/pachirode/monitor/pkg/api/apiserver/v1"
)

func CPUInfoToCPUV1(cpuInfo monitors.CpuInfo) *apiv1.CPU {
	var protoCPU apiv1.CPU
	_ = core.CopyWithConverters(&protoCPU, cpuInfo)
	return &protoCPU
}
