package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// AppConfig drzi konfiguraci aplikace
type AppConfig struct {
	EnvPrefix       string
	Port            int
	ServiceName     string
	JaegerCollector string
	JaegerAgent     string
}

// Init nacte konfiguraci
func (s *AppConfig) Init() error {
	viper.SetConfigType("yaml")
	viper.SetConfigName("config")
	viper.AddConfigPath("config")

	if s.EnvPrefix != "" {
		viper.SetEnvPrefix(s.EnvPrefix)
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	if s.ServiceName == "" {
		MustMapString(&s.ServiceName, "service_name")
	}

	MustMapInt(&s.Port, "port")

	MustMapString(&s.JaegerCollector, "jaeger_collector")
	MustMapString(&s.JaegerAgent, "jaeger_agent")

	return nil
}

// MapString Do cílového pole načte string. Pokud nenajde příslušné nastavení, použije defaultValue
func MapString(target *string, envKey string, defaultValue string) {
	v := viper.GetString(envKey)
	if v == "" {
		v = defaultValue
	}

	*target = v
}

// MapInt Do cílového pole načte int. Pokud nenajde příslušné nastavení, použije defaultValue
func MapInt(target *int, envKey string, defaultValue int) {
	v := viper.GetInt(envKey)
	if v == 0 {
		v = defaultValue
	}
	*target = v
}

// MustMapString Do cílového pole načte string. Pokud nenajde příslušné nastavení, zpanikaří
func MustMapString(target *string, envKey string) {
	v := viper.GetString(envKey)
	if v == "" {
		panic(fmt.Sprintf("Chybi env var: %v", envKey))
	}
	*target = v
}

// MustMapInt Do cílového pole načte int. Pokud nenajde příslušné nastavení, zpanikaří
func MustMapInt(target *int, envKey string) {
	v := viper.GetInt(envKey)
	if v <= 0 {
		panic(fmt.Sprintf("Chybi env var: %v", envKey))
	}
	*target = v
}
