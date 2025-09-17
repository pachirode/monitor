package monitors

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCPUMonitor(t *testing.T) {
	monitor := NewCPUMonitor()
	assert.NotNil(t, monitor)

	println(fmt.Printf("CPU Usage: %.2f%%\n", monitor.getPercent()))
	for _, c := range monitor.getPerCPU() {
		fmt.Printf("%v: user=%v system=%v idle=%v\n",
			c["cpu"], c["user"], c["system"], c["idle"])
	}

	println(monitor.getModel())
}
