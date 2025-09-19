package monitors

import (
	"fmt"

	"github.com/shirou/gopsutil/mem"
)

type MemoryMonitor struct {
	monitor
	MemoryInfo
}

type MemoryInfo struct {
	MemoryAvailable string
	MemoryTotal     string
	MemoryFree      string
	MemoryUsage     string
	MemoryPercent   string
	SwapTotal       string
	SwapFree        string
	SwapUsage       string
	SwapPercent     string
}

func NewMemoryMonitor() *MemoryMonitor {
	return &MemoryMonitor{
		MemoryInfo: MemoryInfo{},
	}
}

func (m *MemoryMonitor) getMemoryStaticInfo() {
	memStat, _ := mem.VirtualMemory()
	swapStat, _ := mem.SwapMemory()

	m.MemoryTotal = fmt.Sprintf("%.2f GB", float64(memStat.Total)/(1024*1024*1024))
	m.SwapTotal = fmt.Sprintf("%.2f GB", float64(swapStat.Total)/(1024*1024*1024))
}

func (m *MemoryMonitor) getMemoryDynamicInfo() {
	memStat, _ := mem.VirtualMemory()
	swapStat, _ := mem.SwapMemory()

	m.MemoryAvailable = fmt.Sprintf("%.2f GB", float64(memStat.Available)/(1024*1024*1024))
	m.MemoryUsage = fmt.Sprintf("%.2f GB", float64(memStat.Used)/(1024*1024*1024))
	m.MemoryFree = fmt.Sprintf("%.2f GB", float64(memStat.Free)/(1024*1024*1024))
	m.MemoryPercent = fmt.Sprintf("%.2f %%", memStat.UsedPercent)
	m.SwapUsage = fmt.Sprintf("%.2f GB", float64(swapStat.Used)/(1024*1024*1024))
	m.SwapFree = fmt.Sprintf("%.2f GB", float64(swapStat.Free)/(1024*1024*1024))
	m.SwapPercent = fmt.Sprintf("%.2f %%", swapStat.UsedPercent)
}

func (m *MemoryMonitor) Update() {
	m.getMemoryStaticInfo()
	m.getMemoryDynamicInfo()
}

func (m *MemoryMonitor) GetMemoryInfo() MemoryInfo {
	m.Update()
	return m.MemoryInfo
}
