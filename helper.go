package main

import (
	"os"
	"time"

	procmeminfo "github.com/guillermo/go.procmeminfo"
	"github.com/matishsiao/goInfo"

	wire "myntapm/glowroot-go-agent/org_glowroot_wire_api_model"
)

func getTimeNowUnixNanoInt64() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond) / int64(time.Nanosecond)
}

func getEnvironment() *wire.Environment {

	var env *wire.Environment

	meminfo := &procmeminfo.MemInfo{}
	meminfo.Update()

	gi := goInfo.GetInfo()

	hostInfo := &wire.HostInfo{
		HostName:                 gi.Hostname,
		AvailableProcessors:      int32(gi.CPUs),
		TotalPhysicalMemoryBytes: &wire.OptionalInt64{Value: int64(meminfo.Total())},
		OsName:    gi.OS,
		OsVersion: gi.Core,
	}
	env = &wire.Environment{
		HostInfo: hostInfo,
		ProcessInfo: &wire.ProcessInfo{
			ProcessId: &wire.OptionalInt64{Value: int64(os.Getpid())},
			StartTime: getTimeNowUnixNanoInt64(),
		},
		JavaInfo: &wire.JavaInfo{},
	}

	return env
}

func getAgentConfig() *wire.AgentConfig {

	var gaugeConfig []*wire.AgentConfig_GaugeConfig
	var alertConfig []*wire.AgentConfig_AlertConfig
	var pluginConfig []*wire.AgentConfig_PluginConfig
	var instrumentationConfig []*wire.AgentConfig_InstrumentationConfig

	agentConfig := &wire.AgentConfig{
		AgentVersion: "0.01",
		TransactionConfig: &wire.AgentConfig_TransactionConfig{
			SlowThresholdMillis:     &wire.OptionalInt32{Value: 2000},
			ProfilingIntervalMillis: &wire.OptionalInt32{Value: 1000},
			CaptureThreadStats:      true,
		},
		UiConfig: &wire.AgentConfig_UiConfig{
			DefaultDisplayedPercentile:      []float64{50, 95, 99},
			DefaultDisplayedTransactionType: "Web",
		},
		UserRecordingConfig: &wire.AgentConfig_UserRecordingConfig{
			User: []string{},
			ProfilingIntervalMillis: &wire.OptionalInt32{Value: 1000},
		},
		AdvancedConfig: &wire.AgentConfig_AdvancedConfig{
			WeavingTimer:                          false,
			ImmediatePartialStoreThresholdSeconds: &wire.OptionalInt32{Value: 60},
			MaxAggregateQueriesPerType:            &wire.OptionalInt32{Value: 500},
			MaxAggregateServiceCallsPerType:       &wire.OptionalInt32{Value: 500},
			MaxAggregateTransactionsPerType:       &wire.OptionalInt32{Value: 500},
			MaxStackTraceSamplesPerTransaction:    &wire.OptionalInt32{Value: 50000},
			MaxTraceEntriesPerTransaction:         &wire.OptionalInt32{Value: 2000},
			MbeanGaugeNotFoundDelaySeconds:        &wire.OptionalInt32{Value: 60},
		},
		GaugeConfig:           gaugeConfig,
		AlertConfig:           alertConfig,
		PluginConfig:          pluginConfig,
		InstrumentationConfig: instrumentationConfig,
	}

	return agentConfig
}
