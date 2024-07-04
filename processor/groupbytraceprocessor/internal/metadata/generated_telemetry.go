// Code generated by mdatagen. DO NOT EDIT.

package metadata

import (
	"errors"

	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/metric/noop"
	"go.opentelemetry.io/otel/trace"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config/configtelemetry"
)

func Meter(settings component.TelemetrySettings) metric.Meter {
	return settings.MeterProvider.Meter("otelcol/groupbytrace")
}

func Tracer(settings component.TelemetrySettings) trace.Tracer {
	return settings.TracerProvider.Tracer("otelcol/groupbytrace")
}

// TelemetryBuilder provides an interface for components to report telemetry
// as defined in metadata and user config.
type TelemetryBuilder struct {
	meter                                   metric.Meter
	ProcessorGroupbytraceConfNumTraces      metric.Int64Gauge
	ProcessorGroupbytraceEventLatency       metric.Int64Histogram
	ProcessorGroupbytraceIncompleteReleases metric.Int64Counter
	ProcessorGroupbytraceNumEventsInQueue   metric.Int64Gauge
	ProcessorGroupbytraceNumTracesInMemory  metric.Int64Gauge
	ProcessorGroupbytraceSpansReleased      metric.Int64Counter
	ProcessorGroupbytraceTracesEvicted      metric.Int64Counter
	ProcessorGroupbytraceTracesReleased     metric.Int64Counter
	level                                   configtelemetry.Level
}

// telemetryBuilderOption applies changes to default builder.
type telemetryBuilderOption func(*TelemetryBuilder)

// WithLevel sets the current telemetry level for the component.
func WithLevel(lvl configtelemetry.Level) telemetryBuilderOption {
	return func(builder *TelemetryBuilder) {
		builder.level = lvl
	}
}

// NewTelemetryBuilder provides a struct with methods to update all internal telemetry
// for a component
func NewTelemetryBuilder(settings component.TelemetrySettings, options ...telemetryBuilderOption) (*TelemetryBuilder, error) {
	builder := TelemetryBuilder{level: configtelemetry.LevelBasic}
	for _, op := range options {
		op(&builder)
	}
	var err, errs error
	if builder.level >= configtelemetry.LevelBasic {
		builder.meter = Meter(settings)
	} else {
		builder.meter = noop.Meter{}
	}
	builder.ProcessorGroupbytraceConfNumTraces, err = builder.meter.Int64Gauge(
		"processor_groupbytrace_conf_num_traces",
		metric.WithDescription("Maximum number of traces to hold in the internal storage"),
		metric.WithUnit("1"),
	)
	errs = errors.Join(errs, err)
	builder.ProcessorGroupbytraceEventLatency, err = builder.meter.Int64Histogram(
		"processor_groupbytrace_event_latency",
		metric.WithDescription("How long the queue events are taking to be processed"),
		metric.WithUnit("ms"), metric.WithExplicitBucketBoundaries([]float64{5, 10, 20, 50, 100, 200, 500, 1000, 2000, 5000}...),
	)
	errs = errors.Join(errs, err)
	builder.ProcessorGroupbytraceIncompleteReleases, err = builder.meter.Int64Counter(
		"processor_groupbytrace_incomplete_releases",
		metric.WithDescription("Releases that are suspected to have been incomplete"),
		metric.WithUnit("<nil>"),
	)
	errs = errors.Join(errs, err)
	builder.ProcessorGroupbytraceNumEventsInQueue, err = builder.meter.Int64Gauge(
		"processor_groupbytrace_num_events_in_queue",
		metric.WithDescription("Number of events currently in the queue"),
		metric.WithUnit("1"),
	)
	errs = errors.Join(errs, err)
	builder.ProcessorGroupbytraceNumTracesInMemory, err = builder.meter.Int64Gauge(
		"processor_groupbytrace_num_traces_in_memory",
		metric.WithDescription("Number of traces currently in the in-memory storage"),
		metric.WithUnit("1"),
	)
	errs = errors.Join(errs, err)
	builder.ProcessorGroupbytraceSpansReleased, err = builder.meter.Int64Counter(
		"processor_groupbytrace_spans_released",
		metric.WithDescription("Spans released to the next consumer"),
		metric.WithUnit("1"),
	)
	errs = errors.Join(errs, err)
	builder.ProcessorGroupbytraceTracesEvicted, err = builder.meter.Int64Counter(
		"processor_groupbytrace_traces_evicted",
		metric.WithDescription("Traces evicted from the internal buffer"),
		metric.WithUnit("1"),
	)
	errs = errors.Join(errs, err)
	builder.ProcessorGroupbytraceTracesReleased, err = builder.meter.Int64Counter(
		"processor_groupbytrace_traces_released",
		metric.WithDescription("Traces released to the next consumer"),
		metric.WithUnit("1"),
	)
	errs = errors.Join(errs, err)
	return &builder, errs
}
