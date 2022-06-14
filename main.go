package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/realbucksavage/simple-user-service/generated/users"
	"google.golang.org/grpc"
)

func main() {

	var (
		grpcPort = flag.Int("grpc-port", 0, "Specify the port number on which the gRPC server will listen. 0 or no value means any random free port.")
	)
	flag.Parse()

	grpcServer := grpc.NewServer()
	users.RegisterUserServiceServer(grpcServer, &defaultUserService{usersMap: userRegistry})

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *grpcPort))
	if err != nil {
		log.Fatalf("cannot create listener for grpc server: %v", err)
	}

	log.Printf("starting gRPC server on %s", lis.Addr())

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("cannot start gRPC server: %v", err)
	}
}
