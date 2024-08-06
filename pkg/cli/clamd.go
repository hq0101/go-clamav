package cli

import (
	"fmt"
	"strings"
)

type ClamdStats struct {
	Pools       int
	State       string
	ThreadsLive int
	ThreadsIdle int
	ThreadsMax  int
	IdleTimeout int
	QueueItems  int
	Stats       float64
	Heap        float64
	Mmap        float64
	Used        float64
	Free        float64
	Releasable  float64
	PoolsUsed   float64
	PoolsTotal  float64
}

func ParsePoolStats(input string) (*ClamdStats, error) {
	lines := strings.Split(input, "\n")
	ps := &ClamdStats{}

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		switch {
		case strings.HasPrefix(line, "POOLS:"):
			fmt.Sscanf(line, "POOLS: %d", &ps.Pools)
		case strings.HasPrefix(line, "STATE:"):
			stateParts := strings.Fields(line)
			if len(stateParts) >= 2 {
				ps.State = stateParts[1]
			}
		case strings.HasPrefix(line, "THREADS:"):
			var extra string
			fmt.Sscanf(line, "THREADS: live %d idle %d max %d idle-timeout %d%s",
				&ps.ThreadsLive, &ps.ThreadsIdle, &ps.ThreadsMax, &ps.IdleTimeout, &extra)
			if strings.Contains(extra, "PRIMARY") {
				ps.State += " PRIMARY"
			}
		case strings.HasPrefix(line, "QUEUE:"):
			fmt.Sscanf(line, "QUEUE: %d items", &ps.QueueItems)
		case strings.HasPrefix(line, "STATS"):
			fmt.Sscanf(line, "STATS %f", &ps.Stats)
		case strings.HasPrefix(line, "MEMSTATS:"):
			fmt.Sscanf(line, "MEMSTATS: heap %fM mmap %fM used %fM free %fM releasable %fM pools %*d pools_used %fM pools_total %fM",
				&ps.Heap, &ps.Mmap, &ps.Used, &ps.Free, &ps.Releasable, &ps.PoolsUsed, &ps.PoolsTotal)
		}
	}
	return ps, nil
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
