package api

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"time"
)

var (
	cfg *Config
)

type Config struct {
	Listen           string        `mapstructure:"listen"`
	ClamdAddress     string        `mapstructure:"clamd_address"`
	ClamdNetworkType string        `mapstructure:"clamd_network_type"`
	ClamdConnTimeout time.Duration `mapstructure:"clamd_conn_timeout"`
	ClamdReadTimeout time.Duration `mapstructure:"clamd_read_timeout"`
}

func Init(configPath string) error {
	viper.SetConfigFile(configPath)
	if _err := viper.ReadInConfig(); _err != nil {
		return fmt.Errorf("failed to read the configuration file: %w", _err)
	}
	if _err := viper.Unmarshal(&cfg); _err != nil {
		return fmt.Errorf("failed to parse the configuration file: %w", _err)
	}
	viper.OnConfigChange(func(e fsnotify.Event) {
		if _err := viper.Unmarshal(&cfg); _err != nil {
			fmt.Println("Failed to reload the configuration file:", _err)
		} else {
			fmt.Println("reload config file")
		}
	})
	viper.WatchConfig()
	return nil
}

func GetCfg() *Config {
	return cfg
}
