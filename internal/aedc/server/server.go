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
	pb.UnimplementedAEDCServiceServer
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
	a.grpcServer.RegisterService(&pb.AEDCService_ServiceDesc, a)
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

func (a *API) VerifyInfo(ctx context.Context, req *pb.VerifyInfoReq) (*pb.VerifyInfoResp, error) {
	// TODO: Get number from kafka stream and do some authentication
	log.Printf("full_name: %s ; number: %s", "Sample Name", req.Number)

	return &pb.VerifyInfoResp{
		FullName: "Sample Name",
		Status:   true,
	}, nil
}
