package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func LoadConfig() AppConfig {
	var config AppConfig

	viper.SetDefault("battery.path", "/sys/class/power_supply/BAT0/capacity")
	viper.SetDefault("battery.low_level_limit", 30)
	viper.SetDefault("battery.high_level_limit", 80)
	viper.SetDefault("charger.path", "/sys/class/power_supply/AC/online")
	viper.SetDefault("sleep_time_in_seconds", 60)

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	viper.AddConfigPath("$HOME/.config/battery-watcher")

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("ℹ️ Nenhum arquivo de config encontrado, usando valores padrão")
	} else {
		fmt.Println("📄 Configuração carregada:", viper.ConfigFileUsed())
	}

	err = viper.Unmarshal(&config)

	return config
}
