package visigrpc

import (
	"github.com/getsentry/raven-go"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type grpcError struct {
	code codes.Code
	err  error
}

func (ge *grpcError) isInternalServerError() bool {
	return ge.code == codes.Unknown || ge.code == codes.Internal || ge.code == codes.DataLoss
}

func (ge *grpcError) log() {
	raven.CaptureError(ge.err, nil)
}

func NewGrpcError(code codes.Code, err error) error {
	ge := grpcError{code: code, err: err}
	if ge.isInternalServerError() {
		ge.log()
	}

	return status.Error(code, err.Error())
}
