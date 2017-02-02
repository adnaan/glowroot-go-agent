package main

import (
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"

	wire "myntapm/glowroot-go-agent/org_glowroot_wire_api_model"

	"github.com/kr/pretty"
)

type GRPCClient struct {
	conn            *grpc.ClientConn
	host            string
	agentID         string
	agentRollupID   string
	collectorClient wire.CollectorServiceClient
	streamClient    wire.CollectorService_CollectAggregateStreamClient
}

func main() {

	// initializing client struct
	grpcClient, err := getGRPCClient()
	if err != nil {
		log.Println(err.Error())
		return
	}
	// defer grpcClient.close()

	initMessage, err := grpcClient.getInitMessage()
	if err != nil {
		log.Println("here:", err.Error())
		return
	}
	pretty.Println(initMessage)

	grpcClient.collectAggregateStream()
}

func (grpcClient *GRPCClient) getInitMessage() (*wire.InitResponse, error) {

	initMessage := &wire.InitMessage{
		AgentId:       grpcClient.agentID,
		AgentRollupId: grpcClient.agentRollupID,
		Environment:   getEnvironment(),
		AgentConfig:   getAgentConfig(),
	}

	initResponse, err := grpcClient.collectorClient.CollectInit(context.Background(), initMessage)
	if err != nil {
		return nil, err
	}

	return initResponse, nil

}

func (grpcClient *GRPCClient) collectAggregateStream() {
	// sending stream header
	aggregateStreamHeader := &wire.AggregateStreamHeader{
		AgentId:     grpcClient.agentID,
		CaptureTime: getTimeNowUnixNanoInt64(),
	}
	aggregateStreamMessage := &wire.AggregateStreamMessage{
		Message: &wire.AggregateStreamMessage_Header{Header: aggregateStreamHeader},
	}
	grpcClient.send(aggregateStreamMessage)

	// sending shared query text
	aggregateSharedQueryText := &wire.Aggregate_SharedQueryText{
		FullText:      "",
		TruncatedText: "",
		FullTextSha1:  "",
	}
	aggregateStreamMessage = &wire.AggregateStreamMessage{
		Message: &wire.AggregateStreamMessage_SharedQueryText{
			SharedQueryText: aggregateSharedQueryText,
		},
	}
	grpcClient.send(aggregateStreamMessage)

	// dummy aggregate
	aggregate := &wire.Aggregate{
		ErrorCount:         2,
		TransactionCount:   90,
		TotalDurationNanos: 3000,
	}

	// transaction aggregate
	transactionAggregate := &wire.TransactionAggregate{
		TransactionType: "Web",
		TransactionName: "/api/test1/test23",
		Aggregate:       aggregate,
	}
	aggregateStreamMessage = &wire.AggregateStreamMessage{
		Message: &wire.AggregateStreamMessage_TransactionAggregate{
			TransactionAggregate: transactionAggregate,
		},
	}
	grpcClient.send(aggregateStreamMessage)

	// overall transaction aggregate
	overallAggregate := &wire.OverallAggregate{
		TransactionType: "Web",
		Aggregate:       aggregate,
	}
	aggregateStreamMessage = &wire.AggregateStreamMessage{
		Message: &wire.AggregateStreamMessage_OverallAggregate{
			OverallAggregate: overallAggregate,
		},
	}
}

func (grpcClient *GRPCClient) send(msg *wire.AggregateStreamMessage) error {
	if err := grpcClient.streamClient.Send(msg); err != nil {
		grpclog.Printf("%v.Send(%v) = ERROR(%v)", grpcClient.streamClient, msg, err)
		return err
	}
	return nil
}

func (grpcClient *GRPCClient) close() {

	reply, err := grpcClient.streamClient.CloseAndRecv()
	if err != nil {
		grpclog.Fatalf("%v.CloseAndRecv() got error %v, want %v", grpcClient.streamClient, err, nil)
	}

	grpclog.Printf("Stream reply: %v", reply)

	err = grpcClient.conn.Close()
	if err != nil {
		grpclog.Fatalf("%v.CloseAndRecv() got error %v, want %v", grpcClient.streamClient, err, nil)
	}

}

func getGRPCClient() (*GRPCClient, error) {

	var err error

	grpcClient := GRPCClient{
		host:          "0.0.0.0:8181",
		agentID:       "echoservice-nmc-1",
		agentRollupID: "echoservice",
	}

	// establishing connection
	grpcClient.conn, err = grpc.Dial(grpcClient.host, grpc.WithInsecure())
	if err != nil {
		log.Printf("failed to connect: %s", err)
		return nil, err
	}

	// initializing collector client
	grpcClient.collectorClient = wire.NewCollectorServiceClient(grpcClient.conn)

	// initializing stream client
	grpcClient.streamClient, err = grpcClient.collectorClient.CollectAggregateStream(context.Background())
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return &grpcClient, nil
}
