package conversion

import (
	"github.com/pachirode/pkg/core"

	"github.com/pachirode/monitor/internal/apiserver/pkg/monitors"
	apiv1 "github.com/pachirode/monitor/pkg/api/apiserver/v1"
)

func MemoryInfoToMemoryV1(info monitors.MemoryInfo) *apiv1.Memory {
	var protoMemory apiv1.Memory
	_ = core.CopyWithConverters(&protoMemory, info)
	return &protoMemory
}
