package visigrpc

import (
	"github.com/getsentry/raven-go"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
)

type grpcError struct {
	Code codes.Code
	Err  error
}

func (ge *grpcError) isInternalServerError() bool {
	return ge.Code == codes.Unknown || ge.Code == codes.Internal || ge.Code == codes.DataLoss
}

func (ge *grpcError) log() {
	raven.CaptureError(ge.Err, nil)
}

func Error(code codes.Code, err error) error {
	ge := grpcError{Code: code, Err: err}
	if ge.isInternalServerError() {
		ge.log()
	}

	return status.Error(code, err.Error())
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
