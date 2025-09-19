package monitors

import (
	"fmt"
	"log"

	"github.com/shirou/gopsutil/disk"
)

type DiskMonitor struct {
	monitor
	DiskInfos map[string]DiskInfo
}

type DiskInfo struct {
	MountPoint string
	Fstype     string
	TotalSize  string
	UsedSize   string
	FreeSize   string
	Usage      string
}

func NewDiskMonitor() *DiskMonitor {
	return &DiskMonitor{
		DiskInfos: make(map[string]DiskInfo),
	}
}

func (m *DiskMonitor) getDiskStaticInfo() {
	partitions, err := disk.Partitions(false)
	if err != nil {
		log.Println("Error getting disk partitions:", err)
	}

	for _, partition := range partitions {
		usage, err := disk.Usage(partition.Mountpoint)
		if err != nil {
			log.Println("Error getting disk usage for", partition.Mountpoint, err)
			continue
		}

		m.DiskInfos[partition.Device] = DiskInfo{
			MountPoint: partition.Mountpoint,
			Fstype:     partition.Fstype,
			TotalSize:  fmt.Sprintf("%.2f GB", float64(usage.Total)/(1024*1024*1024)),
		}
	}
}

func (m *DiskMonitor) getDiskDynamicInfo() {
	partitions, err := disk.Partitions(false)
	if err != nil {
		log.Println("Error getting disk partitions:", err)
	}

	for _, partition := range partitions {
		usage, err := disk.Usage(partition.Mountpoint)
		if err != nil {
			log.Println("Error getting disk usage for", partition.Mountpoint, err)
			continue
		}

		diskInfo := m.DiskInfos[partition.Device]
		diskInfo.UsedSize = fmt.Sprintf("%.2f GB", float64(usage.Used)/(1024*1024*1024))
		diskInfo.FreeSize = fmt.Sprintf("%.2f GB", float64(usage.Free)/(1024*1024*1024))
		diskInfo.Usage = fmt.Sprintf("%.2f %%", usage.UsedPercent)
		m.DiskInfos[partition.Device] = diskInfo
	}
}

func (m *DiskMonitor) Update() {
	m.getDiskStaticInfo()
	m.getDiskDynamicInfo()
}

func (m *DiskMonitor) GetDiskInfos() map[string]DiskInfo {
	m.Update()
	return m.DiskInfos
}
