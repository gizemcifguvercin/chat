package main

import (
	"context"
	"flag"
	"github.com/gizemcifguvercin/chat/log"
	"github.com/gizemcifguvercin/chat/proto"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"

	"google.golang.org/grpc"
)

const (
	defaultName = "john"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address")
	name = flag.String("name", defaultName, "Name")
)

func main() {
	flag.Parse()
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(logger.UnaryInterceptor))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := proto.NewChatClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SayHello(ctx, &proto.ChatRequest{Name: *name, Texts: []string{"hi", "there"}})
	if err != nil {
		log.Fatalf("could not chat: %v", err)
	}
	log.Printf("server response : %s", r.GetMessage())
}

/*
go run main.go --name=Gizem
*/

/*
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    proto/chat.proto
*/
