package visigrpc

import (
	"context"
	"errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

func PanicInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	defer func() {
		rec := recover()
		if rec != nil {
			var e error

			switch t := rec.(type) {
			case string:
				e = errors.New(t)
			case error:
				e = t
			default:
				e = errors.New(`unknown error`)
			}

			err = NewGrpcError(codes.Internal, e)
		}
	}()

	return handler(ctx, req)
}
