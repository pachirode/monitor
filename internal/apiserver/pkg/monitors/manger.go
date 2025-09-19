package monitors

type MonitorManager struct {
	*CpuMonitor
	*DiskMonitor
	*MemoryMonitor
	*NetworkMonitor
}

func NewMonitorManager() *MonitorManager {
	return &MonitorManager{
		CpuMonitor:     NewCpuMonitor(),
		DiskMonitor:    NewDiskMonitor(),
		MemoryMonitor:  NewMemoryMonitor(),
		NetworkMonitor: NewNetworkMonitor(),
	}
}
