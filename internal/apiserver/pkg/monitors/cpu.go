package monitors

import (
	"github.com/shirou/gopsutil/v3/cpu"
)

type CPUMonitor struct {
	monitor
	info          cpu.InfoStat
	percent       float64
	perCPUPercent []map[string]interface{}
}

func NewCPUMonitor() *CPUMonitor {
	info, _ := cpu.Info()
	var cpuInfo cpu.InfoStat
	if len(info) > 0 {
		cpuInfo = info[0]
	}

	return &CPUMonitor{
		monitor: monitor{
			staticStats:  map[string]interface{}{},
			dynamicStats: map[string]interface{}{},
		},
		info:          cpuInfo,
		percent:       0,
		perCPUPercent: []map[string]interface{}{},
	}
}

func (m *CPUMonitor) getPercent() float64 {
	percent, _ := cpu.Percent(0, false)
	if len(percent) > 0 {
		m.percent = percent[0]
	}
	return m.percent
}

func (m *CPUMonitor) getPerCPU() []map[string]interface{} {
	m.perCPUPercent = []map[string]interface{}{}
	percent, _ := cpu.Times(true)

	for _, t := range percent {
		cpuMap := map[string]interface{}{
			"cpu":    t.CPU,
			"user":   t.User,
			"system": t.System,
			"idle":   t.Idle,
		}
		m.perCPUPercent = append(m.perCPUPercent, cpuMap)
	}
	return m.perCPUPercent
}

func (m *CPUMonitor) getModel() string {
	return m.info.ModelName
}
