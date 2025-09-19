package conversion

import (
	"github.com/pachirode/pkg/core"

	"github.com/pachirode/monitor/internal/apiserver/pkg/monitors"
	apiv1 "github.com/pachirode/monitor/pkg/api/apiserver/v1"
)

func DiskInfoToDiskV1(diskInfos map[string]monitors.DiskInfo) []*apiv1.Disk {
	var diskList = make([]*apiv1.Disk, 0)
	for _, diskInfo := range diskInfos {
		var protoCPU apiv1.Disk
		_ = core.CopyWithConverters(&protoCPU, diskInfo)
		diskList = append(diskList, &protoCPU)
	}
	return diskList
}
