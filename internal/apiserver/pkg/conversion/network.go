package conversion

import (
	"github.com/pachirode/pkg/core"

	"github.com/pachirode/monitor/internal/apiserver/pkg/monitors"
	apiv1 "github.com/pachirode/monitor/pkg/api/apiserver/v1"
)

func NetworkInfoToNetworkV1(infos map[string]monitors.NetworkInfo) []*apiv1.Network {
	var networkList = make([]*apiv1.Network, 0)
	for _, networkInfo := range infos {
		var protoCPU apiv1.Network
		_ = core.CopyWithConverters(&protoCPU, networkInfo)
		networkList = append(networkList, &protoCPU)
	}
	return networkList
}
