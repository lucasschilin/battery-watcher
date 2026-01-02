package config

type AppConfig struct {
	Battery BatteryConfig `mapstructure:"battery"`
	Charger ChargerConfig `mapstructure:"charger"`

	SleepTimeInSeconds int `mapstructure:"sleep_time_in_seconds"`
}

type BatteryConfig struct {
	Path           string `mapstructure:"path"`
	LowLevelLimit  int    `mapstructure:"low_level_limit"`
	HighLevelLimit int    `mapstructure:"high_level_limit"`
}

type ChargerConfig struct {
	Path string `mapstructure:"path"`
}
