package monitors

type MonitorManager struct {
	*CpuMonitor
	*DiskMonitor
}

func NewMonitorManager() *MonitorManager {
	return &MonitorManager{
		CpuMonitor:  NewCpuMonitor(),
		DiskMonitor: NewDiskMonitor(),
	}
}
