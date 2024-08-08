package cli

import (
	"fmt"
	"strconv"
	"strings"
)

type ClamdStats struct {
	Pools         int
	ThreadsLive   int
	ThreadsIdle   int
	ThreadsMax    int
	IdleTimeout   int
	QueueItems    int
	MemHeap       float64
	MemMmap       float64
	MemUsed       float64
	MemFree       float64
	MemReleasable float64
	MemPoolsUsed  float64
	MemPoolsTotal float64
}

func ParseStatStr(statStr string) *ClamdStats {
	lines := strings.Split(statStr, "\n")
	stats := &ClamdStats{}

	for _, line := range lines {
		parts := strings.Fields(line)
		if len(parts) == 0 {
			continue
		}
		switch parts[0] {
		case "POOLS:":
			stats.Pools, _ = strconv.Atoi(parts[1])
		case "THREADS:":
			stats.ThreadsLive, _ = strconv.Atoi(parts[2])
			stats.ThreadsIdle, _ = strconv.Atoi(parts[4])
			stats.ThreadsMax, _ = strconv.Atoi(parts[6])
			stats.IdleTimeout, _ = strconv.Atoi(parts[8])
		case "QUEUE:":
			stats.QueueItems, _ = strconv.Atoi(parts[1])
		case "MEMSTATS:":
			stats.MemHeap, _ = strconv.ParseFloat(strings.TrimSuffix(parts[2], "M"), 64)
			stats.MemMmap, _ = strconv.ParseFloat(strings.TrimSuffix(parts[4], "M"), 64)
			stats.MemUsed, _ = strconv.ParseFloat(strings.TrimSuffix(parts[6], "M"), 64)
			stats.MemFree, _ = strconv.ParseFloat(strings.TrimSuffix(parts[8], "M"), 64)
			stats.MemReleasable, _ = strconv.ParseFloat(strings.TrimSuffix(parts[10], "M"), 64)
			stats.MemPoolsUsed, _ = strconv.ParseFloat(strings.TrimSuffix(parts[14], "M"), 64)
			stats.MemPoolsTotal, _ = strconv.ParseFloat(strings.TrimSuffix(parts[16], "M"), 64)
		}
	}

	return stats
}

type ScanResult struct {
	Path   string `json:"path"`
	Virus  string `json:"virus"`
	Status string `json:"status"`
}

func FormatScanResult(result string) ([]ScanResult, error) {
	var scanResults []ScanResult
	lines := strings.Split(result, "\n")

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		parts := strings.SplitN(line, ": ", 3)
		if len(parts) < 2 {
			return nil, fmt.Errorf("invalid result format: %s", line)
		}

		path := parts[0]
		statusPart := parts[1]

		var virus string
		var status string

		if strings.Contains(statusPart, "FOUND") {
			// Extract virus name and status
			virusParts := strings.SplitN(statusPart, " ", 2)
			if len(virusParts) == 2 {
				virus = virusParts[0]
				status = virusParts[1]
			} else {
				status = statusPart
			}
		} else {
			status = statusPart
		}

		scanResults = append(scanResults, ScanResult{
			Path:   path,
			Virus:  virus,
			Status: status,
		})
	}

	return scanResults, nil
}
