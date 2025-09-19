package monitors

import (
	"fmt"
	"os/exec"
	"runtime"
	"strings"

	"github.com/StackExchange/wmi"
	"github.com/pachirode/pkg/log"
	"github.com/shirou/gopsutil/v3/cpu"
)

type CpuMonitor struct {
	monitor
	info          cpu.InfoStat
	percent       float64
	perCPUPercent []map[string]interface{}
}

type CpuInfo struct {
	Model            string
	Architecture     string
	PhysicalCores    int
	LogicalCpus      int
	CurrentFrequency string
	Usage            string
}

func NewCpuMonitor() *CpuMonitor {
	info, _ := cpu.Info()
	var cpuInfo cpu.InfoStat
	if len(info) > 0 {
		cpuInfo = info[0]
	}

	return &CpuMonitor{
		monitor: monitor{
			StaticStats:  map[string]interface{}{},
			DynamicStats: map[string]interface{}{},
		},
		info:          cpuInfo,
		percent:       0,
		perCPUPercent: []map[string]interface{}{},
	}
}

func (m *CpuMonitor) getPercent() float64 {
	percent, _ := cpu.Percent(0, false)
	if len(percent) > 0 {
		m.percent = percent[0]
	}
	return m.percent
}

func (m *CpuMonitor) getPerCPU() []map[string]interface{} {
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

func (m *CpuMonitor) getModel() string {
	return m.info.ModelName
}

func (m *CpuMonitor) getCurrentFrequency() (float64, error) {
	switch runtime.GOOS {
	case "windows":
		return getCurrentCPUFrequencyWindows()
	case "linux":
		return getCurrentCPUFrequencyLinux()
	case "darwin":
		return getCurrentCPUFrequencyMacOS()

	default:
		return 0, fmt.Errorf("%s not support get CPU frequency", runtime.GOOS)
	}
}

type Win32Processor struct {
	CurrentClockSpeed uint32 // 以 MHz 为单位
}

func getCurrentCPUFrequencyWindows() (float64, error) {
	// 查询 WMI 中的 Win32_Processor 类
	var dst []Win32Processor
	query := `SELECT CurrentClockSpeed FROM Win32_Processor`
	err := wmi.Query(query, &dst)
	if err != nil {
		return 0, err
	}

	if len(dst) > 0 {
		return float64(dst[0].CurrentClockSpeed) / 1000, nil
	}
	return 0, fmt.Errorf("current CPU frequency not found")
}

func getCurrentCPUFrequencyLinux() (float64, error) {
	cmd := exec.Command("lscpu")
	out, err := cmd.Output()
	if err != nil {
		return 0, err
	}

	// 解析输出，查找 CPU MHz 信息
	for _, line := range strings.Split(string(out), "\n") {
		if strings.Contains(line, "CPU MHz") {
			parts := strings.SplitN(line, ":", 2)
			if len(parts) == 2 {
				var freq float64
				_, err := fmt.Sscanf(parts[1], "%f", &freq)
				if err != nil {
					return 0, err
				}
				return freq / 1000, nil // 转换为 GHz
			}
		}
	}
	return 0, fmt.Errorf("current CPU frequency not found")
}

func getCurrentCPUFrequencyMacOS() (float64, error) {
	cmd := exec.Command("sysctl", "-n", "hw.cpufrequency")
	out, err := cmd.Output()
	if err != nil {
		return 0, err
	}

	// 将频率值转换为 GHz
	freqMHz := string(out)
	var freq float64
	_, err = fmt.Sscanf(freqMHz, "%f", &freq)
	if err != nil {
		return 0, err
	}

	return freq / 1000000, nil // 转换为 GHz
}

func (m *CpuMonitor) Update() {
	cores, _ := cpu.Counts(false)
	logical, _ := cpu.Counts(true)

	m.monitor.StaticStats = map[string]interface{}{
		"model":          m.getModel(),
		"architecture":   runtime.GOARCH,
		"physical_cores": cores,
		"logical_cpus":   logical,
	}

	usage := m.getPercent()
	m.monitor.DynamicStats = map[string]interface{}{
		"usage": fmt.Sprintf("%.1f %%", usage),
	}

	freq, err := m.getCurrentFrequency()
	if err == nil {
		m.monitor.DynamicStats["current_frequency"] = fmt.Sprintf("%.1f GHz", freq)
	} else {
		log.Fatalf("Get current CPU frequency failed: %v", err)
		m.monitor.DynamicStats["current_frequency"] = fmt.Sprintf("%.1f GHz", freq)
	}
}

func (m *CpuMonitor) GetCpuInfo() CpuInfo {
	m.Update()
	return CpuInfo{
		Model:            m.monitor.StaticStats["model"].(string),
		Architecture:     m.monitor.StaticStats["architecture"].(string),
		PhysicalCores:    m.monitor.StaticStats["physical_cores"].(int),
		LogicalCpus:      m.monitor.StaticStats["logical_cpus"].(int),
		Usage:            m.monitor.DynamicStats["usage"].(string),
		CurrentFrequency: m.monitor.DynamicStats["current_frequency"].(string),
	}
}
