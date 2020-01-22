package visigrpc

import (
	"errors"
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

	return Error(codes.Unknown, e)
}
