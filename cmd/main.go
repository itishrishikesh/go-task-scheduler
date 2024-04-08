package main

import (
	"context"
	"github.com/itishrishikesh/go-task-scheduler/messages"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

func main() {
	conn, err := grpc.Dial("localhost:8082", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := messages.NewTaskManagerClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.ScheduleTask(ctx, &messages.Task{
		Name: "JugMug",
	})
	if err != nil {
		log.Fatalf("could not schedule: %v", err)
	}
	log.Printf("Scheduled task: %v", r)
}
