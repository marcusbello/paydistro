package server

import (
	"context"
	"github.com/marcusbello/paydistro/pkg/kafka"
	pb "github.com/marcusbello/paydistro/proto/v1"
	"google.golang.org/grpc"
	"log"
	"net"
	"sync"
)

type API struct {
	pb.UnimplementedTokenServiceServer
	addr       string
	mu         sync.Mutex
	grpcServer *grpc.Server
	kafka      *kafka.Kafka
}

func New(addr, kafkaURL, topic, groupID string) (*API, error) {
	var opts []grpc.ServerOption
	k, err := kafka.New(kafkaURL, topic)
	if err != nil {
		return nil, err
	}
	a := &API{
		addr:       addr,
		grpcServer: grpc.NewServer(opts...),
		kafka:      k,
	}
	a.grpcServer.RegisterService(&pb.TokenService_ServiceDesc, a)
	return a, nil
}

func (a *API) Start() error {
	a.mu.Lock()
	defer a.mu.Unlock()

	lis, err := net.Listen("tcp", a.addr)
	if err != nil {
		return err
	}

	return a.grpcServer.Serve(lis)
}

func (a *API) GetUser(ctx context.Context, req *pb.GetUserReq) (*pb.GetUserResp, error) {
	// TODO: 1. Send getuserfromaedc event to kafka to `buy_token` topic, get back the user details
	message := struct {
		key   string
		value map[string]interface{}
	}{
		key: req.Number,
		value: map[string]interface{}{
			"number": req.Number,
			"amount": req.Amount,
		},
	}
	err := a.kafka.Writer(ctx, message.key, message.value)
	if err != nil {
		return nil, err
	}

	//// TODO: Get full_name from kafka stream based on number
	//fullNameChannel := make(chan string, 1)
	//errChannel := make(chan error, 1)
	//go func() {
	//	msg, err := a.kafka.Reader()
	//	if err != nil {
	//		errChannel <- err
	//	}
	//	fullName := msg["full_name"]
	//	fullNameChannel <- fullName.(string)
	//}()

	//for {
	//	select {
	//	case result := <-fullNameChannel:
	//		fmt.Println(result)
	//		log.Printf("number: %s ;full_name: %s  amount #{%v}", req.Number, result, req.Amount)
	//		return &pb.GetUserResp{
	//			FullName: result,
	//			Number:   req.Number,
	//		}, nil
	//	case errResult := <-errChannel:
	//		fmt.Println(errResult)
	//		log.Printf("error on getting user with number: %s ;Err: %v", req.Number, errResult)
	//		return &pb.GetUserResp{
	//			FullName: "",
	//			Number:   "",
	//		}, errResult
	//	}
	//}
	//

	log.Printf("number: %s ;full_name: 'Sample Name';  amount #{%v}", req.Number, req.Amount)
	return &pb.GetUserResp{
		FullName: "Sample Name",
		Number:   req.Number,
	}, nil
}

// TODO: Add Confirm() it will take everything and generate payment number
