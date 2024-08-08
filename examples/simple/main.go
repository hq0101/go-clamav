package main

import (
	"encoding/json"
	"fmt"
	"github.com/hq0101/go-clamav/pkg/clamav"
	"time"
)

func main() {
	client := clamav.NewClamClient("tcp", "192.168.127.131:3310", 10*time.Second, 30*time.Second)

	response, err := client.Ping()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Response from ClamAV:", response)

	version, err := client.Version()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("ClamAV Version:", version)

	stats, err := client.Stats()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	jsonResults, err := json.MarshalIndent(stats, "", "    ")
	if err != nil {
		fmt.Printf("Failed to marshal results to JSON: %v\n", err)
		return
	}

	fmt.Println("ClamAV Stats:", string(jsonResults))
}
