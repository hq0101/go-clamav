package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/hq0101/go-clamav/pkg/clamav"
	"github.com/hq0101/go-clamav/pkg/cli"
	"strings"
	"time"
)

func createClient(networkType, address string, connTimeout, readTimeout time.Duration) (*clamav.ClamClient, error) {
	switch networkType {
	case cli.TCP.String(), cli.Unix.String():
	default:
		return nil, fmt.Errorf("invalid network type: %s", networkType)
	}
	return clamav.NewClamClient(networkType, address, connTimeout, readTimeout), nil
}

func handleResponse(results string, err error) {
	if err != nil {
		fmt.Printf("Command failed: %v\n", err)
		return
	} else {
		fmt.Println(results)
	}
}

func pretty(out string, results []cli.ScanResult, err error) {
	if err != nil {
		fmt.Printf("Command failed: %v\n", err)
		return
	}
	switch out {
	case cli.Json.String():
		formatJson(results)
	case cli.Text.String():
		formatText(results)
	}
}

func formatJson(results []cli.ScanResult) {
	jsonResults, err := json.MarshalIndent(results, "", "    ")
	if err != nil {
		fmt.Printf("Failed to marshal results to JSON: %v\n", err)
		return
	}

	fmt.Println(string(jsonResults))
}

func formatText(results []cli.ScanResult) {
	var sb strings.Builder
	for _, result := range results {
		sb.WriteString(fmt.Sprintf("%s %s %s\n", result.Path, result.Virus, result.Status))
	}
	fmt.Println(sb.String())
}

func formatPoolStats(ps *cli.ClamdStats) string {
	return fmt.Sprintf(
		`
=== POOL STATISTICS ===
Pools: %d
Primary Threads:
  - Live: %d
  - Idle: %d
  - Max: %d
  - Idle Timeout: %d
Queue: %d items
Memory Stats:
  - Heap: %.3fM
  - Mmap: %.3fM
  - Used: %.3fM
  - Free: %.3fM
  - Releasable: %.3fM
  - Pools Used: %.3fM
  - Pools Total: %.3fM
`, ps.Pools,
		ps.ThreadsLive, ps.ThreadsIdle, ps.ThreadsMax, ps.IdleTimeout,
		ps.QueueItems,
		ps.MemHeap, ps.MemMmap, ps.MemUsed, ps.MemFree, ps.MemReleasable, ps.MemPoolsUsed, ps.MemPoolsTotal)
}
