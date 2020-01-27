package status

import (
	"github.com/getsentry/raven-go"
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
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
	code := codes.Unknown
	msg := "unknown error"

	st, ok := status.FromError(err)
	if ok {
		code = st.Code()
		msg = st.Message()
	}

	return &Status{Code: uint32(code), Message: msg}
}

func GrpcCodeFromHttpStatus(status int) codes.Code {
	switch status {
	case http.StatusOK:
		return codes.OK
	case http.StatusRequestTimeout:
		return codes.Canceled
	case http.StatusBadRequest:
		return codes.InvalidArgument
	case http.StatusGatewayTimeout:
		return codes.DeadlineExceeded
	case http.StatusNotFound:
		return codes.NotFound
	case http.StatusConflict:
		return codes.AlreadyExists
	case http.StatusForbidden:
		return codes.PermissionDenied
	case http.StatusTooManyRequests:
		return codes.ResourceExhausted
	case http.StatusPreconditionFailed:
		return codes.FailedPrecondition
	case http.StatusNotImplemented:
		return codes.Unimplemented
	case http.StatusInternalServerError:
		return codes.Internal
	case http.StatusServiceUnavailable:
		return codes.Unavailable
	case http.StatusUnauthorized:
		return codes.Unauthenticated
	}

	return codes.Unknown
}
