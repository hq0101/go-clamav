package main

import (
	"github.com/hq0101/go-clamav/internal/api"
	"github.com/spf13/cobra"
	"log"
)

func main() {
	var cfgFilePath string
	var rootCmd = &cobra.Command{
		Use:   "clamd",
		Short: "clamd",
		Run: func(cmd *cobra.Command, args []string) {
			if _err := api.Init(cfgFilePath); _err != nil {
				log.Fatalln(_err)
			}
			if err := api.Run(); err != nil {
				log.Fatalln(err)
			}
		},
	}
	rootCmd.PersistentFlags().StringVarP(&cfgFilePath, "config", "c", "./configs/clamav-api.yaml", "config file path")
	if err := rootCmd.Execute(); err != nil {
		log.Fatalln(err)
	}
}
