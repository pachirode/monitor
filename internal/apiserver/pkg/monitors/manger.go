package monitors

type MonitorManager struct {
	*CPUMonitor
}

func NewMonitorManager() *MonitorManager {
	return &MonitorManager{
		CPUMonitor: NewCPUMonitor(),
	}
}
