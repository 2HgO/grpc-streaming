package main

import (
	"context"
	"io/ioutil"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
	pb "videoStream/source"

	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/status"
)

const port = ":55053"

type server struct{}

func (s *server) GetVideo(info *pb.VideoInfo, stream pb.Source_GetVideoServer) error {
	file, err := os.OpenFile("Daffodil - 34826_dashinit.mp4", os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Println(err)
		return err
	}
	defer file.Close()
	wholeVideo, err := ioutil.ReadAll(file)

	if err != nil {
		log.Println(err)
		return err
	}

	manifest := [][2]int{
		{0,824},
		{825,1775313},
		{1775314,3408415},
		{3408416,5052250},
		{5052251,6685746},
		{6685747,8288919},
		{8288920,9871888},
		{9871889,11436236},
		{11436237,12992586},
		{12992587,14618186},
		{14618187,16262450},
		{16262451,17941891},
		{17941892,19583616},
		{19583617,21249101},
		{21249102,22905881},
		{22905882,24626698},
		{24626699,26253554},
		{26253555,27853963},
		{27853964,29435161},
		{29435162,31016544},
		{31016545,31981952},
	}

	for _, mRange := range manifest {
		if err := stream.Send(&pb.Video{Data: wholeVideo[mRange[0]:mRange[1]+1]}); err != nil {
			log.Println("Error sending data: ", err)
			return err
		}
	}
	return nil
}

func main() {
	defer log.Println("Server shutdown successful")
	defer func() {
		if err := recover(); err != nil {
			println()
			log.Fatalln(err)
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
		log.Fatalln(err)
	}	

	s := grpc.NewServer(grpc.UnaryInterceptor(RecoveryInterceptor))
	pb.RegisterSourceServer(s, &server{})
	
	go func (){
		log.Println("Server starting up at 0.0.0.0:55053 ...")
		if err := s.Serve(lis); err != nil {
			errChan <- err
		}
	}()
	defer s.GracefulStop()

	select {
		case <-channel:
			println()
			log.Println("Source Server shutting down...")
			return
		case err := <-errChan:
			panic(err)
	}
}

// RecoveryInterceptor ...
func RecoveryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (res interface{}, err error) {
	defer func() {
		if erro := recover(); erro != nil {
			res, err = nil, status.Errorf(codes.Internal, "%v", erro)
			grpclog.SetLogger(log.New(os.Stdout, "[SourceService::Error] ", 0))
			grpclog.Errorf("Error occured during RPC method=%s; Error=%v", info.FullMethod, erro)
		}
	}()

	start := time.Now()

	resp, erro := handler(ctx, req)

	grpclog.SetLogger(log.New(os.Stdout, "[SourceService::Log] ", 0))
	grpclog.Printf("Handled RPC method=%s; Duration=%s; Error=%v", info.FullMethod, time.Since(start), erro)
	return resp, erro
}