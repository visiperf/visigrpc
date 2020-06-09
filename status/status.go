package status

import (
	"errors"
	"net/http"

	"github.com/getsentry/sentry-go"
	spb "google.golang.org/genproto/googleapis/rpc/status"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func newStatus(code codes.Code, msg string) *status.Status {
	if code == codes.Unknown || code == codes.Internal || code == codes.DataLoss {
		sentry.CaptureException(errors.New(msg))
	}

	return status.New(code, msg)
}

func New(code codes.Code, msg string) *spb.Status {
	return newStatus(code, msg).Proto()
}

func Error(code codes.Code, msg string) error {
	return newStatus(code, msg).Err()
}

func FromError(err error) *spb.Status {
	code := codes.Unknown
	msg := "unknown error"

	st, ok := status.FromError(err)
	if ok {
		code = st.Code()
		msg = st.Message()
	}

	return status.New(code, msg).Proto()
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
