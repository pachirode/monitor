package monitors

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCPUMonitor(t *testing.T) {
	monitor := NewCpuMonitor()
	assert.NotNil(t, monitor)

	println(fmt.Printf("CPU Usage: %.2f%%\n", monitor.getPercent()))
	for _, c := range monitor.getPerCPU() {
		fmt.Printf("%v: user=%v system=%v idle=%v\n",
			c["cpu"], c["user"], c["system"], c["idle"])
	}

	println(monitor.getModel())

	freq, _ := monitor.getCurrentFrequency()
	fmt.Printf("%.1f GHz", freq)
}

func TestMonitor(t *testing.T) {
	m := NewCpuMonitor()
	assert.NotNil(t, m)
	m.Update()

	for k, v := range m.GetStatsInfo() {
		fmt.Printf("%s : %s\n", k, v)
	}

	println(m.GetStatsJSON())
}
