package cmd

import (
	"fmt"
	"github.com/hq0101/go-clamav/pkg/cli"
	"github.com/spf13/cobra"
	"log"
)

func NewStatsCmd(p cli.Params) *cobra.Command {
	return &cobra.Command{
		Use:   "stats",
		Short: "Get ClamAV stats",
		Run: func(cmd *cobra.Command, args []string) {
			client, _err := createClient(p.GetNetworkType(), p.GetAddress(), p.GetConnTimeout(), p.GetReadTimeout())
			if _err != nil {
				log.Fatalln(_err)
			}
			response, err := client.Stats()
			if err != nil {
				log.Fatalln(err)
			}
			fmt.Println(formatPoolStats(response))
		},
	}
}
