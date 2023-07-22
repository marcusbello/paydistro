package client

import (
	"context"
	pb "github.com/marcusbello/paydistro/proto/v1"
	"google.golang.org/grpc"
	"time"
)

type Client struct {
	client pb.TokenServiceClient
	conn   *grpc.ClientConn
}

func New(addr string) (*Client, error) {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	return &Client{
		client: pb.NewTokenServiceClient(conn),
		conn:   conn,
	}, nil
}

func (c *Client) GetUser(ctx context.Context, amount, number string) (string, string, error) {
	if _, ok := ctx.Deadline(); !ok {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, 2*time.Second)
		defer cancel()
	}
	resp, err := c.client.GetUser(ctx, &pb.GetUserReq{
		Number: number,
		Amount: amount,
	})
	if err != nil {
		return "", "", err
	}
	return resp.FullName, resp.Number, nil
}
