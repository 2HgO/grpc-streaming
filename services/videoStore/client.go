package main

import (
	"log"
	"time"

	pb "videoStream/auth"
	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

var kacp = keepalive.ClientParameters{
	Time:                10 * time.Second,
	Timeout:             time.Second,
	PermitWithoutStream: true,
}

var UserClient = func() pb.AuthClient {
	conn, err := grpc.Dial("auth-server:55056", grpc.WithInsecure(), grpc.WithUserAgent("source-service/0.1"), grpc.WithKeepaliveParams(kacp))
	if err != nil {
		log.Fatalln(err)
	}
	return pb.NewAuthClient(conn)
}()
