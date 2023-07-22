package server

import (
	"context"
	pb "github.com/marcusbello/paydistro/proto/v1"
	"google.golang.org/grpc"
	"log"
	"net"
	"sync"
)

type API struct {
	pb.UnimplementedWalletServiceServer
	addr       string
	mu         sync.Mutex
	grpcServer *grpc.Server
}

func New(addr string) (*API, error) {
	var opts []grpc.ServerOption
	a := &API{
		addr:       addr,
		grpcServer: grpc.NewServer(opts...),
	}
	a.grpcServer.RegisterService(&pb.WalletService_ServiceDesc, a)
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

func (a *API) SendPayment(ctx context.Context, req *pb.SendPaymentReq) (*pb.SendPaymentResp, error) {
	// TODO: get amount and number from the kafka stream and number, send money to acctNumber
	log.Printf("status: %v ;accountNumber: %s ;amount : %s", false, req.AcctNumber, req.Amount)
	//  TODO: some magic e.g id acctNum, remove your %, send the rest to aedc, must be successful, send token
	return &pb.SendPaymentResp{
		Token:  "1234-5678-9012-3456-7890",
		Status: true,
	}, nil
}
