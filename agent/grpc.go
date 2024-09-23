package main

import (
	"context"
	"log"

	"google.golang.org/grpc"
)

func JobCall(key string) bool {

	req := JobRequest{
		Key: key,
	}

	resp, err := (*Ctx.JobService).CallJob(context.Background(), &req, grpc.WaitForReady(true))
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("success: %v", resp.Ret)
	ret := false
	if resp.Ret == 1 {
		ret = true
	}
	return ret
}

func connectToGrpc() *grpc.ClientConn {
	// Replace the following with your gRPC server details
	host := "[::1]"
	port := "9999"

	// Create a new gRPC client
	conn, err := grpc.Dial(host+":"+port, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatal(err)
	}

	return conn
}
