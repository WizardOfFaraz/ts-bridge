package tsbridge

import (
	"time"

	"github.com/google/ts-bridge/stackdriver"
)

// Config holds all the ConfigOptions
type Config struct {
	Options      ConfigOptions
	Dependencies ExternalDependencies
}

// ConfigOptions  a set of global options required to initialize configuration.
type ConfigOptions struct {
	CounterResetInterval     time.Duration
	EnableStatusPage         bool
	Filename                 string
	MinPointAge              time.Duration
	MonitoringBackends       []string
	SDInternalMetricsProject string
	SDLookBackInterval       time.Duration
	StorageEngine            string
	BoltdbPath               string
	DatastoreProject         string
	UpdateParallelism        int
	UpdateTimeout            time.Duration
	SyncPeriod               time.Duration
}

type ExternalDependencies struct {
	SDClient       *stackdriver.Adapter
	StatsCollector *StatsCollector
}

// NewConfig returns a new ConfigOptions struct.
func NewConfig(options *ConfigOptions) *Config {
	return &Config{
		Options: *options,
	}
}
