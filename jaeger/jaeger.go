package jaeger

import (
	"contrib.go.opencensus.io/exporter/jaeger"
	"go.opencensus.io/trace"

	cfg "github.com/endevelcz/odin/config"
)

// InitJaeger Inicializuje OpenCensus tracing.
func InitJaeger(cfg *cfg.AppConfig) (*jaeger.Exporter, error) {
	exporter, err := jaeger.NewExporter(jaeger.Options{
		AgentEndpoint: cfg.JaegerAgent,
		Process: jaeger.Process{
			ServiceName: cfg.ServiceName,
		},
	})
	if err != nil {
		return nil, err
	}

	trace.RegisterExporter(exporter)
	trace.ApplyConfig(trace.Config{
		DefaultSampler: trace.AlwaysSample(),
	})

	return exporter, nil
}
