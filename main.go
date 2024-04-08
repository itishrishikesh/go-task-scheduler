package main

import (
	"context"
	"github.com/itishrishikesh/go-task-scheduler/messages"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct {
	messages.UnimplementedTaskManagerServer
}

func (s *server) ScheduleTask(ctx context.Context, in *messages.Task) (*messages.Empty, error) {
	log.Println("Schedule Task " + in.Name)
	return &messages.Empty{}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":8082")
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()
	messages.RegisterTaskManagerServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		panic(err)
	}
}
