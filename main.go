package main

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/adnaan/glowroot-go-agent/hdrhist"
	wire "github.com/adnaan/glowroot-go-agent/org_glowroot_wire_api_model"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

type GRPCClient struct {
	conn          *grpc.ClientConn
	host          string
	agentID       string
	agentRollupID string
	client        wire.CollectorServiceClient
	stream        Streams
}

type Streams struct {
	aggregate wire.CollectorService_CollectAggregateStreamClient
	trace     wire.CollectorService_CollectTraceStreamClient
}

func main() {

	// initializing client struct
	grpcClient, err := getGRPCClient("echoservice-nmc-1", "echoservice")
	if err != nil {
		log.Println(err.Error())
		return
	}
	defer grpcClient.close()

	_, err = grpcClient.getInitMessage()
	if err != nil {
		log.Println("here:", err.Error())
		return
	}
	grpcClient.collectAggregateStream()
	grpcClient.collectGaugeValue()
	grpcClient.collectLogValues()
	grpcClient.collectTraceStream()
}

func (grpcClient *GRPCClient) getInitMessage() (*wire.InitResponse, error) {

	initMessage := &wire.InitMessage{
		AgentId:       grpcClient.agentID,
		AgentRollupId: grpcClient.agentRollupID,
		Environment:   getEnvironment(),
		AgentConfig:   getAgentConfig(),
	}

	initResponse, err := grpcClient.client.CollectInit(context.Background(), initMessage)
	if err != nil {
		return nil, err
	}

	return initResponse, nil

}

func (grpcClient *GRPCClient) collectAggregateStream() {
	// sending stream header
	grpcClient.sendAggregate(
		&wire.AggregateStreamMessage{
			Message: &wire.AggregateStreamMessage_Header{Header: &wire.AggregateStreamHeader{
				AgentId:     grpcClient.agentID,
				CaptureTime: getTimeNowUnixNanoInt64(),
			}},
		})

	// sending shared query text
	grpcClient.sendAggregate(
		&wire.AggregateStreamMessage{
			Message: &wire.AggregateStreamMessage_SharedQueryText{
				SharedQueryText: &wire.Aggregate_SharedQueryText{
					FullText:      "",
					TruncatedText: "",
					FullTextSha1:  "",
				},
			},
		})

	// dummy aggregate
	aggregate := &wire.Aggregate{
		TransactionCount:   100,
		ErrorCount:         5,
		TotalDurationNanos: 200,
		AsyncTransactions:  true,

		MainThreadStats: &wire.Aggregate_ThreadStats{
			TotalAllocatedBytes: &wire.OptionalDouble{Value: 1000},
			TotalBlockedNanos:   &wire.OptionalDouble{Value: 20000},
			TotalCpuNanos:       &wire.OptionalDouble{Value: 3000},
			TotalWaitedNanos:    &wire.OptionalDouble{Value: 4000},
		},

		MainThreadRootTimer: []*wire.Aggregate_Timer{
			&wire.Aggregate_Timer{
				Name:       "testabc",
				Extended:   false,
				TotalNanos: 100,
				Count:      23,
				ChildTimer: []*wire.Aggregate_Timer{},
			},
		},

		QueriesByType: []*wire.Aggregate_QueriesByType{
			&wire.Aggregate_QueriesByType{
				Type: "SQL",
				Query: []*wire.Aggregate_Query{
					&wire.Aggregate_Query{
						FullText:             "Select * from myntra",
						SharedQueryTextIndex: 0,
						TotalDurationNanos:   0,
						ExecutionCount:       1000,
						TotalRows:            &wire.OptionalInt64{Value: 10},
					},
				},
			},
		},

		ServiceCallsByType: []*wire.Aggregate_ServiceCallsByType{
			&wire.Aggregate_ServiceCallsByType{
				Type: "GET",
				ServiceCall: []*wire.Aggregate_ServiceCall{
					&wire.Aggregate_ServiceCall{
						Text:               "/api/app1/test/getdat",
						TotalDurationNanos: 10,
						ExecutionCount:     40000,
					},
					&wire.Aggregate_ServiceCall{
						Text:               "/api/app2/test/getdat",
						TotalDurationNanos: 10,
						ExecutionCount:     40000,
					},
					&wire.Aggregate_ServiceCall{
						Text:               "/api/app3/test/getdat",
						TotalDurationNanos: 10,
						ExecutionCount:     40000,
					},
					&wire.Aggregate_ServiceCall{
						Text:               "/api/app4/test/getdat",
						TotalDurationNanos: 10,
						ExecutionCount:     40000,
					},
				},
			},
		},
	}
	// hdr histogram of duration
	hist := hdrhist.WithConfig(
		hdrhist.Config{
			LowestDiscernible: 1000,
			HighestTrackable:  2000,
			SigFigs:           5,
		},
	)
	for i := int64(0); i < 10000; i++ {
		rand.Seed(time.Now().Unix())
		hist.Record(rand.Int63n(1000) + 1000)
	}
	encodedBytes := bytes.NewBuffer([]byte{})
	if err := hist.EncodeInto(encodedBytes, 10000); err != nil {
		log.Println(err.Error())
		return
	}
	aggregate.DurationNanosHistogram = &wire.Aggregate_Histogram{
		OrderedRawValue: []int64{50, 60, 70, 40, 50, 60, 70, 80, 80},
		EncodedBytes:    encodedBytes.Bytes(),
	}

	// overall transaction aggregate
	grpcClient.sendAggregate(
		&wire.AggregateStreamMessage{
			Message: &wire.AggregateStreamMessage_OverallAggregate{
				OverallAggregate: &wire.OverallAggregate{
					TransactionType: "Web",
					Aggregate:       aggregate,
				},
			},
		})

	// transaction aggregate
	grpcClient.sendAggregate(
		&wire.AggregateStreamMessage{
			Message: &wire.AggregateStreamMessage_TransactionAggregate{
				TransactionAggregate: &wire.TransactionAggregate{
					TransactionType: "Web",
					TransactionName: fmt.Sprintf("/api/test%d", 1),
					Aggregate:       aggregate,
				},
			},
		})

}

func (grpcClient *GRPCClient) collectGaugeValue() {

	msg, err := grpcClient.client.CollectGaugeValues(context.Background(),
		&wire.GaugeValueMessage{
			AgentId: grpcClient.agentID,
			GaugeValues: []*wire.GaugeValue{
				&wire.GaugeValue{
					GaugeName:   "name",
					CaptureTime: getTimeNowUnixNanoInt64(),
					Value:       float64(20),
					Weight:      int64(10),
				},
			},
		})
	if err != nil {
		log.Println(err.Error())
		return
	}

	log.Println(msg)
}

func (grpcClient *GRPCClient) collectLogValues() {

	msg, err := grpcClient.client.Log(context.Background(),
		&wire.LogMessage{
			AgentId: grpcClient.agentID,
			LogEvent: &wire.LogEvent{
				Timestamp:  time.Now().Unix(),
				Level:      2,
				LoggerName: "test logger",
				Message:    "test msg",
				Throwable:  nil,
			},
		},
	)

	if err != nil {
		log.Println(err.Error())
		return
	}

	log.Println(msg)
}

func (grpcClient *GRPCClient) collectTraceStream() {
	// sending header
	grpcClient.sendTrace(
		&wire.TraceStreamMessage{
			Message: &wire.TraceStreamMessage_Header{
				Header: &wire.TraceStreamHeader{
					AgentId: grpcClient.agentID,
				},
			},
		})
	// seding shared query test
	grpcClient.sendTrace(
		&wire.TraceStreamMessage{
			Message: &wire.TraceStreamMessage_SharedQueryText{
				SharedQueryText: &wire.Trace_SharedQueryText{
					FullText:         "",
					TruncatedText:    "",
					TruncatedEndText: "",
					FullTextSha1:     "",
				},
			},
		})
	// sending trace
	grpcClient.sendTrace(
		&wire.TraceStreamMessage{
			Message: &wire.TraceStreamMessage_Trace{
				Trace: &wire.Trace{
					Id:                "1",
					MainThreadProfile: nil,
					AuxThreadProfile:  nil,
					Update:            false,
					SharedQueryText: []*wire.Trace_SharedQueryText{
						&wire.Trace_SharedQueryText{
							FullText:         "",
							TruncatedText:    "",
							TruncatedEndText: "",
							FullTextSha1:     "",
						},
					},
					Header: &wire.Trace_Header{
						StartTime:                            10,
						CaptureTime:                          10,
						DurationNanos:                        10,
						EntryCount:                           10,
						MainThreadProfileSampleCount:         10,
						AuxThreadProfileSampleCount:          10,
						Partial:                              false,
						Slow:                                 false,
						Async:                                false,
						EntryLimitExceeded:                   false,
						MainThreadProfileSampleLimitExceeded: false,
						AuxThreadProfileSampleLimitExceeded:  false,
						TransactionType:                      "Web",
						TransactionName:                      "/api/app/test",
						Headline:                             "test api",
						User:                                 "pavan",

						MainThreadRootTimer: &wire.Trace_Timer{
							Name:       "",
							Extended:   false,
							TotalNanos: 1000,
							Count:      30,
							Active:     false,
							ChildTimer: []*wire.Trace_Timer{},
						},

						AuxThreadRootTimer: []*wire.Trace_Timer{
							&wire.Trace_Timer{
								Name:       "",
								Extended:   false,
								TotalNanos: 1000,
								Count:      30,
								Active:     false,
								ChildTimer: []*wire.Trace_Timer{},
							},
						},

						AsyncTimer: []*wire.Trace_Timer{
							&wire.Trace_Timer{
								Name:       "",
								Extended:   false,
								TotalNanos: 1000,
								Count:      30,
								Active:     false,
								ChildTimer: []*wire.Trace_Timer{},
							},
						},

						DetailEntry: []*wire.Trace_DetailEntry{
							&wire.Trace_DetailEntry{
								Name: "",
								Value: []*wire.Trace_DetailValue{
									&wire.Trace_DetailValue{},
								},
								ChildEntry: []*wire.Trace_DetailEntry{},
							},
						},

						Error: &wire.Trace_Error{
							Exception: nil,
							Message:   "No error",
						},

						Attribute: []*wire.Trace_Attribute{
							&wire.Trace_Attribute{
								Name:  "asdas",
								Value: []string{},
							},
						},

						MainThreadStats: &wire.Trace_ThreadStats{
							TotalAllocatedBytes: &wire.OptionalInt64{Value: 1000},
							TotalBlockedNanos:   &wire.OptionalInt64{Value: 20000},
							TotalCpuNanos:       &wire.OptionalInt64{Value: 3000},
							TotalWaitedNanos:    &wire.OptionalInt64{Value: 4000},
						},

						AuxThreadStats: &wire.Trace_ThreadStats{
							TotalAllocatedBytes: &wire.OptionalInt64{Value: 1000},
							TotalBlockedNanos:   &wire.OptionalInt64{Value: 20000},
							TotalCpuNanos:       &wire.OptionalInt64{Value: 3000},
							TotalWaitedNanos:    &wire.OptionalInt64{Value: 4000},
						},
					},
				},
			},
		})

}

func (grpcClient *GRPCClient) sendAggregate(msg interface{}) error {
	if err := grpcClient.stream.aggregate.SendMsg(msg); err != nil {
		grpclog.Printf("%v.Send(%v) = ERROR(%v)", grpcClient.stream.trace, msg, err)
		return err
	}
	return nil
}

func (grpcClient *GRPCClient) sendTrace(msg interface{}) error {
	if err := grpcClient.stream.trace.SendMsg(msg); err != nil {
		grpclog.Printf("%v.Send(%v) = ERROR(%v)", grpcClient.stream.trace, msg, err)
		return err
	}
	return nil
}

func (grpcClient *GRPCClient) closeTrace() {

	reply, err := grpcClient.stream.trace.CloseAndRecv()
	if err != nil {
		grpclog.Fatalf("%v.CloseAndRecv() got error %v, want %v", grpcClient.stream.trace, err, nil)
	}
	grpclog.Printf("Stream reply: %v", reply)

}

func (grpcClient *GRPCClient) closeAggregate() {

	reply, err := grpcClient.stream.aggregate.CloseAndRecv()
	if err != nil {
		grpclog.Fatalf("%v.CloseAndRecv() got error %v, want %v", grpcClient.stream.aggregate, err, nil)
	}
	grpclog.Printf("Stream reply: %v", reply)

}

func (grpcClient *GRPCClient) closeConn() {

	err := grpcClient.conn.Close()
	if err != nil {
		grpclog.Fatalf("%v.CloseAndRecv() got error %v, want %v", grpcClient.stream.trace, err, nil)
	}

}

func (grpcClient *GRPCClient) close() {
	grpcClient.closeAggregate()
	grpcClient.closeTrace()
	grpcClient.closeConn()
}

func getGRPCClient(agentID, agentRollupID string) (*GRPCClient, error) {

	var err error

	grpcClient := GRPCClient{
		host:          "0.0.0.0:8181",
		agentID:       agentID,
		agentRollupID: agentRollupID,
	}

	// establishing connection
	grpcClient.conn, err = grpc.Dial(grpcClient.host, grpc.WithInsecure())
	if err != nil {
		log.Printf("failed to connect: %s", err)
		return nil, err
	}

	// initializing collector client
	grpcClient.client = wire.NewCollectorServiceClient(grpcClient.conn)

	// initializing stream client
	grpcClient.stream.aggregate, err = grpcClient.client.CollectAggregateStream(context.Background())
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	grpcClient.stream.trace, err = grpcClient.client.CollectTraceStream(context.Background())
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return &grpcClient, nil
}
