package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// AppConfig drzi konfiguraci aplikace
type AppConfig struct {
	Port             int
	ServiceName      string
	DbURL            string
	DbMaxConnections int
	JaegerCollector  string
	JaegerAgent      string
}

// Init nacte konfiguraci
func (s *AppConfig) Init(envPrefix string, serviceName string) error {
	viper.SetConfigType("yaml")
	viper.SetConfigName("config")
	viper.AddConfigPath("config")
	viper.SetEnvPrefix(envPrefix)
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	mapString(&s.ServiceName, "service_name", serviceName)

	mustMapInt(&s.Port, "port")
	mustMapString(&s.DbURL, "db_url")
	mustMapInt(&s.DbMaxConnections, "db_max_connections")
	mustMapString(&s.JaegerCollector, "jaeger_collector")
	mustMapString(&s.JaegerAgent, "jaeger_agent")

	return nil
}

func mapString(target *string, envKey string, defaultValue string) {
	v := viper.GetString(envKey)
	if v == "" {
		v = defaultValue
	}

	*target = v
}
func mapInt(target *int, envKey string, defaultValue int) {
	v := viper.GetInt(envKey)
	if v == 0 {
		v = defaultValue
	}
	*target = v
}

func mustMapString(target *string, envKey string) {
	v := viper.GetString(envKey)
	if v == "" {
		panic(fmt.Sprintf("Chybi env var: %v", envKey))
	}
	*target = v
}
func mustMapInt(target *int, envKey string) {
	v := viper.GetInt(envKey)
	if v <= 0 {
		panic(fmt.Sprintf("Chybi env var: %v", envKey))
	}
	*target = v
}
