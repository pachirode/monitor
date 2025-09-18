package monitors

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDiskMonitor(t *testing.T) {
	m := NewDiskMonitor()
	assert.NotNil(t, m)

	m.Update()

	for k, v := range m.DiskInfos {
		fmt.Println(k, v)
	}

}
