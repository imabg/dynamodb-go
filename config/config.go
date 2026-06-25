package config

type Config struct {
	AWS_ACCESS_KEY_ID     string `mapstructure:"AWS_ACCESS_KEY_ID"`
	AWS_SECRET_ACCESS_KEY string `mapstructure:"AWS_SECRET_ACCESS_KEY"`
	AWS_REGION          string `mapstructure:"AWS_REGION"`
	SERVER_PORT          string `mapstructure:"SERVER_PORT"`
}