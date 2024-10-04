package config

import "github.com/spf13/viper"

var Config = &AppConfig{}

type AppConfig struct {
	AppName        string `mapstructure:"app_name"`
	AppVersio      string `mapstructure:"app_version"`
	DBName         string `mapstructure:"db_name"`
	DBHost         string `mapstructure:"db_host"`
	DBPort         string `mapstructure:"db_port"`
	DBUser         string `mapstructure:"db_user"`
	DBPassword     string `mapstructure:"db_password"`
	LoggingLevel   string `mapstructure:"logging_level"`
	HttpServerPort string `mapstructure:"http_server_port"`
	ServerTimeout  string `mapstructure:"server_timeout"`
}

func Load() error {
	v := viper.New()
	v.BindEnv("app_name")
	v.BindEnv("app_version")
	v.BindEnv("db_name")
	v.BindEnv("db_host")
	v.BindEnv("db_port")
	v.BindEnv("db_user")
	v.BindEnv("db_password")
	v.SetDefault("logging_level", "info")
	v.SetDefault("http_server_port", "8080")
	v.SetDefault("server_timeout", "10")

	v.AutomaticEnv()
	if err := v.Unmarshal(Config); err != nil {
		return err
	}

	return nil
}
