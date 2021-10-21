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

func HttpStatusFromGrpcCode(code codes.Code) int {
	switch code {
	case codes.OK:
		return 200
	case codes.Canceled:
		return 499
	case codes.Unknown:
		return 500
	case codes.InvalidArgument:
		return 400
	case codes.DeadlineExceeded:
		return 504
	case codes.NotFound:
		return 404
	case codes.AlreadyExists:
		return 409
	case codes.PermissionDenied:
		return 403
	case codes.ResourceExhausted:
		return 429
	case codes.FailedPrecondition:
		return 400
	case codes.Aborted:
		return 409
	case codes.OutOfRange:
		return 400
	case codes.Unimplemented:
		return 501
	case codes.Internal:
		return 500
	case codes.Unavailable:
		return 503
	case codes.DataLoss:
		return 500
	case codes.Unauthenticated:
		return 401
	}

	return http.StatusInternalServerError
}
