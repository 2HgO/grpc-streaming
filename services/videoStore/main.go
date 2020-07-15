package main

import (
	"net"
	"os"
	"os/signal"
	"syscall"

	pb "videoStream/source"

	grpc "google.golang.org/grpc"
)

const port = ":55053"

func main() {
	defer grpcLog.Infoln("[\033[97;45mSourceService\033[0m] Server shutdown successful")
	defer func() {
		if err := recover(); err != nil {
			grpcLog.Fatalf("\n[\033[97;41mSourceService\033[0m] Erro occured=%v\n", err)
		}
	}()

	channel := make(chan os.Signal, 1)
	defer close(channel)
	errChan := make(chan error, 1)
	defer close(errChan)

	signal.Notify(channel, syscall.SIGINT, syscall.SIGTERM)

	lis, err := net.Listen("tcp", port)
	defer lis.Close()
	if err != nil {
		grpcLog.Fatalf("\n[\033[97;45mSourceService\033[0m] Erro occured=%v\n", err)
	}

	s := grpc.NewServer(grpc.UnaryInterceptor(UnaryRecoveryInterceptor), grpc.StreamInterceptor(StreamRecoveryInterceptor))
	pb.RegisterSourceServer(s, &server{})

	go func() {
		grpcLog.Infoln("[\033[97;45mSourceService\033[0m] Server starting up at 0.0.0.0:55053 ...")
		if err := s.Serve(lis); err != nil {
			errChan <- err
		}
	}()
	defer s.GracefulStop()

	select {
	case <-channel:
		grpcLog.Infoln("\n[\033[97;45mSourceService\033[0m] Source Server shutting down...")
		return
	case err := <-errChan:
		panic(err)
	}
}
