package monitors

import (
	"encoding/json"
	"fmt"

	"k8s.io/kube-openapi/pkg/util/sets"
)

type IMonitor interface {
	Update()
	GetStatsInfo() map[string]interface{}
	GetStatsJSON() string
	PrintStatsInfo()
}

type monitor struct {
	Filed        sets.String
	StaticStats  map[string]interface{}
	DynamicStats map[string]interface{}
}

var _ IMonitor = (*monitor)(nil)

func (m *monitor) Update() {

}

func (m *monitor) GetStatsInfo() map[string]interface{} {
	stats := make(map[string]interface{})
	for k, v := range m.StaticStats {
		stats[k] = v
	}
	for k, v := range m.DynamicStats {
		stats[k] = v
	}
	return stats
}

func (m *monitor) GetStatsJSON() string {
	stats := m.GetStatsInfo()
	statsJSON, _ := json.Marshal(stats)
	return string(statsJSON)
}

func (m *monitor) PrintStatsInfo() {
	stats := m.GetStatsInfo()
	for key, value := range stats {
		fmt.Printf("%s: %v\n", key, value)
	}
}

func (m *monitor) FilterStats(stats interface{}) interface{} {
	switch v := stats.(type) {
	case map[string]interface{}:
		filteredStats := make(map[string]interface{})
		for k, val := range v {
			if m.Filed.Has(k) {
				filteredStats[k] = val
			}
		}
		return filteredStats
	case []interface{}:
		filteredList := []interface{}{}
		for _, item := range v {
			filteredList = append(filteredList, m.FilterStats(item))
		}
		return filteredList
	default:
		// If it's neither a map nor a list, return it as is
		return stats
	}
}
