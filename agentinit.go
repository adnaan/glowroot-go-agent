package main

import (
	"context"
	"log"
	"os"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"

	"github.com/guillermo/go.procmeminfo"
	"github.com/matishsiao/goInfo"

	wire "github.com/adnaan/glowroot/org_glowroot_wire_api_model"
	"github.com/kr/pretty"
)

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

func main() {

	agentID := "echoservice-nmc-1"
	agentRollupID := "echoservice"

	initMessage := &wire.InitMessage{
		AgentId:       agentID,
		AgentRollupId: agentRollupID,
		Environment:   getEnvironment(),
		AgentConfig:   getAgentConfig(),
	}

	sc, err := newInsecureStreamClient("0.0.0.0:8181", initMessage)
	if err != nil {
		return
	}

	aggregate := &wire.Aggregate{}

	txAgg := &wire.TransactionAggregate{
		TransactionType: "Web",
		TransactionName: "/api/test1",
		Aggregate:       aggregate,
	}

	aggStreamHeader := &wire.AggregateStreamHeader{
		AgentId:     agentID,
		CaptureTime: getTimeNowUnixNanoInt64(),
	}

	messageHeader := &wire.AggregateStreamMessage{
		Message: &wire.AggregateStreamMessage_Header{Header: aggStreamHeader},
	}

	sc.send(messageHeader)

	aggSharedQueryText := &wire.Aggregate_SharedQueryText{
		FullText:      "",
		TruncatedText: "",
		FullTextSha1:  "",
	}

	msgSharedQueryText := &wire.AggregateStreamMessage{
		Message: &wire.AggregateStreamMessage_SharedQueryText{
			SharedQueryText: aggSharedQueryText,
		},
	}

	sc.send(msgSharedQueryText)

	overAllAgg := &wire.OverallAggregate{
		TransactionType: "Web",
		Aggregate:       aggregate,
	}

	msgTxOverallAgg := &wire.AggregateStreamMessage{
		Message: &wire.AggregateStreamMessage_OverallAggregate{
			OverallAggregate: overAllAgg,
		},
	}

	sc.send(msgTxOverallAgg)

	msgTxAgg := &wire.AggregateStreamMessage{
		Message: &wire.AggregateStreamMessage_TransactionAggregate{
			TransactionAggregate: txAgg,
		},
	}

	sc.send(msgTxAgg)

	sc.close()

}

type streamClient struct {
	stream wire.CollectorService_CollectAggregateStreamClient
}

func newInsecureStreamClient(host string, initMessage *wire.InitMessage) (*streamClient, error) {

	var sc *streamClient

	conn, err := grpc.Dial(host, grpc.WithInsecure())
	if err != nil {
		log.Printf("failed to connect: %s", err)
		return sc, err
	}
	defer conn.Close()

	collectorServiceClient := wire.NewCollectorServiceClient(conn)

	initResponse, err := collectorServiceClient.CollectInit(context.Background(), initMessage)
	if err != nil {
		log.Println("Error", err, initResponse)
		return sc, err
	}

	pretty.Println(initResponse)

	sc.stream, err = collectorServiceClient.CollectAggregateStream(context.Background())
	if err != nil {
		log.Println("Error", err, initResponse)
		return sc, err
	}

	return sc, nil
}

func (s *streamClient) send(msg *wire.AggregateStreamMessage) error {
	if err := s.stream.Send(msg); err != nil {
		grpclog.Printf("%v.Send(%v) = %v", s.stream, msg, err)
		return err
	}
	return nil
}

func (s *streamClient) close() {

	reply, err := s.stream.CloseAndRecv()
	if err != nil {
		grpclog.Fatalf("%v.CloseAndRecv() got error %v, want %v", s.stream, err, nil)
	}

	grpclog.Printf("Stream reply: %v", reply)
}
