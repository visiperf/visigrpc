package status

import "google.golang.org/grpc/codes"

type Status struct {
	Code    uint32
	Message string
}

func (s *Status) isInternal() bool {
	return false
}

func (s *Status) log() {}

func New(code codes.Code, msg string) *Status {
	return nil
}

func Error(code codes.Code, msg string) error {
	return nil
}

func FromError(err error) *Status {
	return nil
}
