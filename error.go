package visigrpc

import (
	"errors"
	"github.com/getsentry/raven-go"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type grpcError struct {
	Code codes.Code
	Err  error
}

func Error(code codes.Code, err error) error {
	if code == codes.Unknown || code == codes.Internal || code == codes.DataLoss {
		raven.CaptureError(err, nil)
	}

	return status.Error(code, err.Error())
}

func FromError(err error) *grpcError {
	code := codes.Unknown
	msg := "unknown error"

	st, ok := status.FromError(err)
	if ok {
		code = st.Code()
		msg = st.Message()
	}

	return &grpcError{Code: code, Err: errors.New(msg)}
}
