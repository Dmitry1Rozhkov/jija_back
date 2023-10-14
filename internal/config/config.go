package config

import (
	"errors"
	"fmt"
	"github.com/spf13/viper"
	"time"
)

const (
	defaultHTTPRWTimeout          = 10 * time.Second
	defaultHTTPMaxHeaderMegabytes = 1
	defaultKeycloakTimeout        = 3 // seconds
	defaultListCount              = 50
	defaultPage                   = 1
)

type (
	Config struct {
		HTTP        HTTPConfig
		IsDebug     bool
		CorsOrigins []string
		DBConfig    DBConfig
	}

	HTTPConfig struct {
		Host               string
		Port               string
		ReadTimeout        time.Duration
		WriteTimeout       time.Duration
		MaxHeaderMegabytes int
	}

	DBConfig struct {
		Path string
	}
)

func Init(configsDir string) (*Config, error) {
	InitDefault()
	var cfg Config

	err := parseYml(configsDir, &cfg)
	if err != nil {
		return nil, err
	}

	//err = parseEnv(configsDir, &cfg)
	//if err != nil {
	//	return nil, err
	//}

	viper.Reset()
	return &cfg, nil
}

func parseYml(configDir string, cfg *Config) error {
	if err := parseConfigFile(configDir+"internal/config", "yaml"); err != nil {
		fmt.Print(err.Error())
		return err
	}

	if err := viper.UnmarshalKey("http", &cfg.HTTP); err != nil {
		return err
	}

	if err := viper.UnmarshalKey("isDebug", &cfg.IsDebug); err != nil {
		return err
	}

	if err := viper.UnmarshalKey("corsOrigins", &cfg.CorsOrigins); err != nil {
		return err
	}

	return nil
}

func parseEnv(configDir string, cfg *Config) error {
	if err := parseConfigFile(configDir, "env"); err != nil {
		return err
	}

	if err := viper.UnmarshalKey("PATH_DB", &cfg.DBConfig.Path); err != nil {
		return err
	}

	return nil
}

func parseConfigFile(folder string, fileType string) error {
	viper.AddConfigPath(folder)
	switch fileType {
	case "yaml":
		viper.SetConfigName("config")
	case "env":
		viper.SetConfigName("App")
		viper.AutomaticEnv()
	default:
		return errors.New("fileType is invalid")
	}

	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	return viper.MergeInConfig()
}

func InitDefault() {
	viper.SetDefault("http.max_header_megabytes", defaultHTTPMaxHeaderMegabytes)
	viper.SetDefault("http.timeouts.read", defaultHTTPRWTimeout)
	viper.SetDefault("http.timeouts.write", defaultHTTPRWTimeout)
	viper.SetDefault("keycloak.timeout", defaultKeycloakTimeout)
	viper.SetDefault("query_params.page", defaultPage)
	viper.SetDefault("query_params.list", defaultListCount)
}
