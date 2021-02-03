package main

import (
	"context"
	"google.golang.org/grpc"
	pb "my.github/grpc-expose-http/http_proto"
	"net"
)

type RestService struct{}

func (r *RestService) Get(ctx context.Context, message *pb.StringMessage) (*pb.StringMessage, error) {
	return &pb.StringMessage{Value: "Get=" + message.Value}, nil
}

func (r *RestService) Post(ctx context.Context, message *pb.StringMessage) (*pb.StringMessage, error) {
	return &pb.StringMessage{Value: "Post=" + message.Value}, nil
}

func main() {
	grpcServer := grpc.NewServer()
	pb.RegisterRestServiceServer(grpcServer, new(RestService))
	l, _ := net.Listen("tcp", ":5000")
	_ = grpcServer.Serve(l)
}
