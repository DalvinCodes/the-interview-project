package main

import (
	"context"
	"interview-client/internal/consumer"
	"log"

	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	ctx := context.Background()
	serviceAddress := "127.0.0.1:8080"
	conn, err := grpc.DialContext(
		ctx,
		serviceAddress,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		log.Fatalln(errors.Wrap(err, "failed to connect to service"))
	}
	consumer := consumer.New(conn)
	consumer.HelloWorld(ctx)
}
