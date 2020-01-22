# gRPC helpers, middlewares ... for Golang

Table of contents
=================

  * [Install](#install)
  * [Usage](#usage)
    * [Error](#error)
    * [From error](#from-error)
    * [gRPC code from HTTP status](#grpc-code-from-http-status)
    * [Panic interceptor](#panic-interceptor)
  * [References](#references)

## Install

Use `go get` to install this package.

    go get github.com/visiperf/visigrpc


## Usage

### Error

The `Error(code codes.Code, err error) error` function is used to return a gRPC error and log it into [Sentry](https://sentry.io).

```go
type server struct { }

func main() {
  // init Sentry config
  if err := raven.SetDSN(...); err != nil {
    ...
  }
  raven.SetEnvironment(...) 
  
  // gRPC server
  lis, err := net.Listen("tcp", ":9090")
	if err != nil {
    ...
	}
  
  s := grpc.NewServer()
  
  RegisterServiceServer(s, &server{})
  
  if err := s.Serve(lis); err != nil {
    ...
	}
}

func (s *server) Get(ctx context.Context, in *GetRequest) (*GetResponse, error) {
  return nil, visigrpc.Error(codes.Unimplemented, errors.New("implement me"))
}
```

##### IMPORTANT : Only `Unknown`, `Internal` and `DataLoss` errors will be reported in Sentry !

### From Error

If you receive a gRPC error (made with visigrpc.Error(...) or with status.Error(...)), you can decode it with `FromError(err error) *grpcError` to retrieve the gRPC code and the message.

```go
type server struct { }

func main() {
  // init Sentry config
  if err := raven.SetDSN(...); err != nil {
    ...
  }
  raven.SetEnvironment(...) 
  
  // gRPC server
  lis, err := net.Listen("tcp", ":9090")
	if err != nil {
    ...
	}
  
  s := grpc.NewServer()
  
  RegisterServiceServer(s, &server{})
  
  if err := s.Serve(lis); err != nil {
    ...
	}
}

func (s *server) Get(ctx context.Context, in *GetRequest) (*GetResponse, error) {
  return nil, visigrpc.Error(codes.Unimplemented, errors.New("implement me"))
}

func (s *server) Create(ctx context.Context, in *CreateRequest) (*CreateResponse, error) {
  resp, err := s.Get(ctx, &GetRequest{})
  if err != nil {
    ge := visigrpc.FromError(err)
    // ge.Code -> codes.Unimplemented
    // ge.Err -> "implement me"
    ...
  }
  
  return nil, visigrpc.Error(codes.Unimplemented, errors.New("implement me"))
}
```

### gRPC code from HTTP status

### Panic interceptor

## References

* Sentry : [github.com/getsentry/raven-go](https://github.com/getsentry/raven-go)