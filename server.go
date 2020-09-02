package main

import (
	freightQuoter "github.com/guilhermesiani/go-grpc/freight_quoter"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("Failed to listen on port 9000: %v", err)
	}

	s := freightQuoter.Server{}

	grpcServer := grpc.NewServer()

	freightQuoter.RegisterFreightQuoterServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server over port 9000: %v", err)
	}

}
