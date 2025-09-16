package monitors

var (
	CPUMonitor  = "CPU-Monitor"
	DISKMonitor = "DISK-Monitor"
)

type IMonitor interface {
}

type monitor struct {
	monitorType  string
	staticStats  []string
	dynamicStats []string
}

var _ IMonitor = (*monitor)(nil)
