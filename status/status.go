package status

import (
	"github.com/getsentry/raven-go"
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Status struct {
	Code    uint32
	Message string
}

func (s *Status) toError() error {
	return status.Error(codes.Code(s.Code), s.Message)
}

func New(code codes.Code, msg string) *Status {
	if code == codes.Unknown || code == codes.Internal || code == codes.DataLoss {
		raven.CaptureError(errors.New(msg), nil)
	}

	return &Status{Code: uint32(code), Message: msg}
}

func Error(code codes.Code, msg string) error {
	return New(code, msg).toError()
}

func FromError(err error) *Status {
	return nil
}
