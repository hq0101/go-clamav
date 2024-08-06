package cmd

import (
	"github.com/hq0101/go-clamav/pkg/cli"
	"github.com/spf13/cobra"
	"log"
)

func NewShutdownCmd(p cli.Params) *cobra.Command {
	return &cobra.Command{
		Use:   "shutdown",
		Short: "Shutdown the ClamAV server",
		Run: func(cmd *cobra.Command, args []string) {
			client, _err := createClient(p.GetNetworkType(), p.GetAddress(), p.GetConnTimeout(), p.GetReadTimeout())
			if _err != nil {
				log.Fatalln(_err)
			}
			response, err := client.Shutdown()
			handleResponse(response, err)
		},
	}
}
