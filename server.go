package main

import (
	"context"
	"log"
	"net"

	pb "github.com/itok01/grpc-gateway-chat/grpcgatewaychat"
	"google.golang.org/grpc"
)

const (
	port = ":8080"
)

type server struct{}

var messageCount int
var messages = make(map[int]*pb.Message)

func (s *server) GetChat(_ *pb.Null, stream pb.Chat_GetChatServer) error {
	for i := 0; i < messageCount; i++ {
		stream.Send(messages[i])
	}
	return nil
}

func (s *server) PostChat(ctx context.Context, msg *pb.Message) (*pb.MessageOk, error) {
	messages[messageCount] = msg
	messageCount++
	return &pb.MessageOk{Ok: true}, nil
}

func grpcServer() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err)
	}
	s := grpc.NewServer()
	pb.RegisterChatServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
