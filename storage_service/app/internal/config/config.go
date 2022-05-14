package config

import "github.com/spf13/viper"

type Config struct {
	ServerAddr     string `mapstructure:"SS_SERVER_ADDR"`
	ServerHeader   string `mapstructure:"SS_SERVER_HEADER"`
	MinioEndpoint  string `mapstructure:"SS_MINIO_ENDPOINT"`
	MinioSecretKey string `mapstructure:"SS_MINIO_SECRET_KEY"`
	MinioAccessKey string `mapstructure:"SS_MINIO_ACCESS_KEY"`
	MinioUseSSL    bool   `mapstructure:"SS_MINIO_USE_SSL"`
	IsDevelopment  bool   `mapstructure:"SS_IS_DEVELOPMENT"`
}

func NewConfig() (*Config, error) {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	viper.AutomaticEnv()

	config := &Config{
		ServerHeader:  "Socia Network",
		ServerAddr:    ":10000",
		IsDevelopment: true,
	}
	err = viper.Unmarshal(config)
	return config, err
}
