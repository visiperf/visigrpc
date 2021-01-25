package visigrpc

import (
	"crypto/tls"
	"crypto/x509"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

// Dialer is interface implemented by types that create a connection to a gRPC server
type Dialer interface {
	Dial(target string, options ...grpc.DialOption) (*grpc.ClientConn, error)
}

type securedDialer struct{}

// NewSecuredDialer is factory to create a new gRPC dialer using TLS
func NewSecuredDialer() Dialer {
	return &securedDialer{}
}

func (d *securedDialer) Dial(target string, options ...grpc.DialOption) (*grpc.ClientConn, error) {
	cert, err := x509.SystemCertPool()
	if err != nil {
		return nil, err
	}

	cred := credentials.NewTLS(&tls.Config{
		RootCAs: cert,
	})

	return grpc.Dial(target, append(options, grpc.WithAuthority(target), grpc.WithTransportCredentials(cred))...)
}

type insecuredDialer struct{}

// NewInsecuredDialer is factory to create a new gRPC dialer using insecure dial option
func NewInsecuredDialer() Dialer {
	return &insecuredDialer{}
}

func (d *insecuredDialer) Dial(target string, options ...grpc.DialOption) (*grpc.ClientConn, error) {
	return grpc.Dial(target, append(options, grpc.WithInsecure())...)
}
