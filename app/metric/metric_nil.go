package metric

import (
	"github.com/rcrowley/go-metrics"
)

type metricNilImplement struct {
	config MetricConfig
}

func NewNilMetric(config MetricConfig) (Metric, error) {
	return &metricNilImplement{
		config: config,
	}, nil
}

func (this *metricNilImplement) GetCounter(name string) metrics.Counter {
	return metrics.NilCounter{}
}

func (this *metricNilImplement) GetGauge(name string) metrics.Gauge {
	return metrics.NilGauge{}
}

func (this *metricNilImplement) GetGaugeFloat64(name string) metrics.GaugeFloat64 {
	return metrics.NilGaugeFloat64{}
}

func (this *metricNilImplement) GetHistogram(name string) metrics.Histogram {
	return metrics.NilHistogram{}
}

func (this *metricNilImplement) GetMeter(name string) metrics.Meter {
	return metrics.NilMeter{}
}

func (this *metricNilImplement) GetTimer(name string) metrics.Timer {
	return metrics.NilTimer{}
}

func (this *metricNilImplement) Run() error {
	return nil
}

func (this *metricNilImplement) Close() {
}
