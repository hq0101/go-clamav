package cmd

import (
	"github.com/hq0101/go-clamav/pkg/cli"
	"github.com/spf13/cobra"
	"log"
)

func NewMultiScanCmd(p cli.Params) *cobra.Command {
	return &cobra.Command{
		Use:   "multiscan",
		Short: "Multithreaded scan of a file or directory",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			client, _err := createClient(p.GetNetworkType(), p.GetAddress(), p.GetConnTimeout(), p.GetReadTimeout())
			if _err != nil {
				log.Fatalln(_err)
			}
			response, err := client.MultiScan(args[0])
			pretty(p.GetOut().String(), response, err)
		},
	}
}
