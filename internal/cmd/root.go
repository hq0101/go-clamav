package cmd

import (
	"fmt"
	"github.com/hq0101/go-clamav/pkg/cli"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"net"
	"time"
)

const (
	address     = "address"
	ConnTimeout = "conn_timeout"
	netType     = "nettype"
	config      = "config"
	out         = "out"
	ReadTimeout = "read_timeout"
)

func Root(p cli.Params) *cobra.Command {
	var rootCmd = &cobra.Command{
		Use:   "clamd-cli",
		Short: "A CLI for interacting with ClamAV",
		CompletionOptions: cobra.CompletionOptions{
			DisableDefaultCmd: true,
		},
		SilenceErrors: true,
		SilenceUsage:  true,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			return initParams(p, cmd)
		},
	}

	rootCmd.PersistentFlags().StringP(config, "f", "", "config file clamav-cli.yaml)")
	rootCmd.PersistentFlags().StringP(address, "a", "", "ClamAV server address /var/run/clamav/clamd.ctl or 127.0.0.1:3310")
	rootCmd.PersistentFlags().DurationP(ConnTimeout, "t", 10*time.Second, "Connection timeout")
	rootCmd.PersistentFlags().DurationP(ReadTimeout, "r", 30*time.Second, "Read timeout")
	rootCmd.PersistentFlags().StringP(netType, "n", "", "Network type (unix/tcp)")
	rootCmd.PersistentFlags().StringP(out, "o", "text", "json、text (default text)")

	rootCmd.AddCommand(NewPingCmd(p))
	rootCmd.AddCommand(NewVersionCmd(p))
	rootCmd.AddCommand(NewContScanCmd(p))
	rootCmd.AddCommand(NewScanCmd(p))
	rootCmd.AddCommand(NewInstreamCmd(p))
	rootCmd.AddCommand(NewMultiScanCmd(p))
	rootCmd.AddCommand(NewStatsCmd(p))
	rootCmd.AddCommand(NewReloadCmd(p))
	rootCmd.AddCommand(NewShutdownCmd(p))
	rootCmd.AddCommand(NewVersionCommandsCmd(p))
	return rootCmd
}

func initParams(p cli.Params, cmd *cobra.Command) error {
	cfg, err := cmd.Flags().GetString(config)
	if err != nil {
		return err
	}

	if _err := initConfig(cfg, p); _err != nil {
		return _err
	}
	if _err := initializeFlags(cmd, p); _err != nil {
		return _err
	}
	if _err := validateConfig(p); _err != nil {
		return _err
	}
	return nil
}

func initConfig(cfg string, p cli.Params) error {
	if cfg == "" {
		return nil
	}
	viper.SetConfigFile(cfg)
	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("error reading config file: %v", err)
	}

	viper.AutomaticEnv()
	if viper.IsSet("clamd_address") {
		p.SetAddress(viper.GetString("clamd_address"))
	}
	if viper.IsSet("clamd_conn_timeout") {
		p.SetConnTimeout(viper.GetDuration("clamd_conn_timeout"))
	}
	if viper.IsSet("clamd_read_timeout") {
		p.SetReadTimeout(viper.GetDuration("clamd_read_timeout"))
	}
	if viper.IsSet("clamd_network_type") {
		p.SetNetworkType(viper.GetString("clamd_network_type"))
	}
	if viper.IsSet("clamd_out") {
		p.SetOut(cli.OutType(viper.GetString("clamd_out")))
	}
	return nil
}

func initializeFlags(rootCmd *cobra.Command, p cli.Params) error {
	addr, err := rootCmd.Flags().GetString(address)
	if err != nil {
		return err
	}
	if addr != "" && p.GetAddress() == "" {
		p.SetAddress(addr)
	}
	connTimeout, err := rootCmd.Flags().GetDuration(ConnTimeout)
	if err != nil {
		return err
	}
	if connTimeout != 0 && p.GetConnTimeout() == 0 {
		p.SetConnTimeout(connTimeout)
	}
	readTimeout, err := rootCmd.Flags().GetDuration(ReadTimeout)
	if err != nil {
		return err
	}
	if readTimeout != 0 && p.GetReadTimeout() == 0 {
		p.SetReadTimeout(readTimeout)
	}

	netWorkType, err := rootCmd.Flags().GetString(netType)
	if err != nil {
		return err
	}

	if netWorkType != "" && p.GetNetworkType() == "" {
		p.SetNetworkType(netWorkType)
	}

	ot, err := rootCmd.Flags().GetString(out)
	if err != nil {
		return err
	}

	if ot != "" && p.GetOut().String() == "" {
		p.SetOut(cli.OutType(ot))
	}

	return nil
}

func validateConfig(p cli.Params) error {
	if p.GetAddress() == "" {
		return fmt.Errorf("ClamAV server address cannot be empty")
	}
	if p.GetNetworkType() == "" {
		return fmt.Errorf("Network type cannot be empty")
	}

	switch p.GetNetworkType() {
	case cli.Unix.String(), cli.TCP.String():
		if p.GetNetworkType() == cli.TCP.String() {
			if err := validateTCPAddress(p.GetAddress()); err != nil {
				return fmt.Errorf("invalid TCP address format: %v", err)
			}
		}
	default:
		return fmt.Errorf("Invalid network type: %s. Must be 'unix' or 'tcp'", p.GetNetworkType())
	}
	return nil
}

func validateTCPAddress(addr string) error {
	host, port, err := net.SplitHostPort(addr)
	if err != nil {
		return err
	}
	if ip := net.ParseIP(host); ip == nil {
		return fmt.Errorf("invalid IP address")
	}
	if portNum, err := net.LookupPort(cli.TCP.String(), port); err != nil || portNum <= 0 {
		return fmt.Errorf("invalid port number")
	}
	return nil
}
