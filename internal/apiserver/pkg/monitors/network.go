package monitors

import (
	"fmt"
	"log"
	"net"
	"strings"

	gpsNet "github.com/shirou/gopsutil/v3/net"
)

type NetworkMonitor struct {
	monitor
	NetworkInfos map[string]NetworkInfo
}

type NetworkInfo struct {
	Name         string
	BytesSent    string
	BytesReceive string
	Ip           string
	Ipv6         string
}

func NewNetworkMonitor() *NetworkMonitor {
	return &NetworkMonitor{
		NetworkInfos: make(map[string]NetworkInfo, 0),
	}
}

func (m *NetworkMonitor) getNetworkStaticInfo() {
	interfaces, err := net.Interfaces()
	if err != nil {
		log.Fatalf("Error fetching network addresses: %v", err)
		return
	}

	for _, inter := range interfaces {
		addrs, _ := inter.Addrs()
		networkInfo := NetworkInfo{
			Name: inter.Name,
		}
		for _, addr := range addrs {
			ipNet, ok := addr.(*net.IPNet)
			if ok {
				ip := ipNet.IP.String()
				if strings.Contains(ip, ":") {
					networkInfo.Ipv6 = ip
				} else {
					networkInfo.Ip = ip
				}
			}
		}
		m.NetworkInfos[inter.Name] = networkInfo
	}
}

func (m *NetworkMonitor) getNetworkDynamicInfo() {
	netStats, err := gpsNet.IOCounters(true)
	if err != nil {
		log.Fatalf("Error fetching network stats: %v", err)
		return
	}

	for _, stat := range netStats {
		networkInfo := m.NetworkInfos[stat.Name]
		networkInfo.BytesSent = fmt.Sprintf("%.2f MB\n", float64(stat.BytesSent)/1024/1024)
		networkInfo.BytesReceive = fmt.Sprintf("%.2f MB\n", float64(stat.BytesRecv)/1024/1024)
		m.NetworkInfos[stat.Name] = networkInfo
	}
}

func (m *NetworkMonitor) Update() {
	m.getNetworkStaticInfo()
	m.getNetworkDynamicInfo()
}

func (m *NetworkMonitor) GetNetworkInfos() map[string]NetworkInfo {
	m.Update()
	return m.NetworkInfos
}
