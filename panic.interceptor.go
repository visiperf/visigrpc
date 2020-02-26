package visigrpc

import (
	"errors"
	"github.com/visiperf/visigrpc/v2/status"
	"google.golang.org/grpc/codes"
)

func RecoveryHandler(p interface{}) error {
	var e error

	switch t := p.(type) {
	case string:
		e = errors.New(t)
	case error:
		e = t
	default:
		e = errors.New(`unknown error`)
	}

	return status.Error(codes.Unknown, e.Error())
}
