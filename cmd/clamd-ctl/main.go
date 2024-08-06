package main

import (
	"fmt"
	"github.com/hq0101/go-clamav/internal/cmd"
	"github.com/hq0101/go-clamav/pkg/cli"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

func searchSubCommand(rootCmd *cobra.Command, args []string) *cobra.Command {
	for _, _cmd := range rootCmd.Commands() {
		if len(args) < 1 {
			return nil
		}
		if strings.HasPrefix(args[0], _cmd.Name()) {
			return _cmd
		}
	}
	return nil
}

func main() {
	p := &cli.ClamdParams{}
	rootCmd := cmd.Root(p)
	ag := os.Args[1:]
	subcommand := searchSubCommand(rootCmd, ag)
	if subcommand == nil && len(ag) > 0 {
		_ = rootCmd.Help()
		os.Exit(1)
	}
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
