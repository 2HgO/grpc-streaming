package main

import (
	"context"
	"os"
	"time"
	"videoStream/auth"

	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

var grpcLog grpclog.LoggerV2

func init() {
	grpcLog = grpclog.NewLoggerV2(os.Stdout, os.Stderr, os.Stderr)
	grpclog.SetLoggerV2(grpcLog)
}

const (
	UnaryInit string = "[\033[97;44mSourceService\033[0m] Handling RPC method=%s; Origin=%s\n"
	UnaryInfo string = "[%sSourceService\033[0m] Handled RPC method=%s; Duration=%s; Error=%v"
	UnaryError string = "[\033[97;41mSourceService\033[0m] error occured during RPC method=%s; Error=%v"

	StreamInit string = "[\033[97;44mSourceService\033[0m] Handling RPC stream method=%s; Origin=%s\n"
	StreamInfo string = "[%sSourceService\033[0m] Handled RPC stream method=%s; Duration=%s; Error=%v"
	StreamError string = "[\033[97;41mSourceService\033[0m] error occured during RPC stream method=%s; Error=%v"
)

// UnaryRecoveryInterceptor is a unary interceptor for the Source service and is used to recover from unexpected panics as well as log requests to the server
func UnaryRecoveryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (res interface{}, err error) {
	defer func() {
		if erro := recover(); erro != nil {
			res, err = nil, status.Errorf(codes.Internal, "%v", erro)
			grpcLog.Errorf(UnaryError, info.FullMethod, erro)
		}
	}()
	origin := "UNKNOWN"
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Error(codes.InvalidArgument, "no metadata")
	}
	orMD := md.Get("origin")
	if len(orMD) > 0 {
		origin = orMD[0]
	}
	auMd := md.Get("authorization")
	if len(auMd) < 1 {
		return nil, status.Error(codes.Unauthenticated, "authentication unsuccessful")
	}
	token := auMd[0]

	if res, rErr := UserClient.ValidateToken(context.TODO(), &auth.ValidationInfo{Token: token}); rErr != nil || !res.Success {
		return nil, status.Error(codes.Unauthenticated, "authentication unsuccessful")
	}

	grpcLog.Infof(UnaryInit, info.FullMethod, origin)
	start := time.Now()

	resp, erro := handler(ctx, req)

	grpcLog.Infof(UnaryInfo, getColor(erro), info.FullMethod, origin, time.Since(start), erro)
	return resp, erro
}

// StreamRecoveryInterceptor is a stream interceptor for the Source service and is used to recover from unexpected panics as well as log requests to the server
func StreamRecoveryInterceptor(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) (err error) {
	defer func() {
		if erro := recover(); erro != nil {
			err = status.Errorf(codes.Internal, "%v", erro)
			grpcLog.Errorf(StreamError, info.FullMethod, erro)
		}
	}()
	origin := "UNKNOWN"
	md, ok := metadata.FromIncomingContext(ss.Context())
	if !ok {
		return status.Error(codes.InvalidArgument, "no metadata")
	}
	orMD := md.Get("origin")
	if len(orMD) > 0 {
		origin = orMD[0]
	}
	auMd := md.Get("authorization")
	if len(auMd) < 1 {
		return status.Error(codes.Unauthenticated, "authentication unsuccessful")
	}
	token := auMd[0]

	if res, rErr := UserClient.ValidateToken(context.TODO(), &auth.ValidationInfo{Token: token}); rErr != nil || !res.Success {
		return status.Error(codes.Unauthenticated, "authentication unsuccessful")
	}
	grpcLog.Infof(StreamInit, info.FullMethod, origin)
	start := time.Now()

	err = handler(srv, ss)

	grpcLog.Infof(StreamInfo, getColor(err), info.FullMethod, origin, time.Since(start), err)
	return err
}

func getColor(err error) string {
	if err == nil {
		return "\033[97;42m"
	}
	return "\033[90;43m"
}
