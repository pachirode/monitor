package monitors

type IMonitor interface {
}

type monitor struct {
	staticStats  map[string]interface{}
	dynamicStats map[string]interface{}
}

var _ IMonitor = (*monitor)(nil)
