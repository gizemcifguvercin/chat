package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/gizemcifguvercin/chat/proto"
	"log"
	"net"

	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type server struct {
	proto.UnimplementedChatServer
}

func (s *server) SayHello(ctx context.Context, in *proto.ChatRequest) (*proto.ChatReply, error) {
	log.Printf("client request: %v %v", in.GetName(), in.GetTexts())
	return &proto.ChatReply{Message: "Hello " + in.GetName()}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	proto.RegisterChatServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
