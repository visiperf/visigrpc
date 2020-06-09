# gRPC helpers, middlewares ... for Golang

Package `visigrpc` provide some functions to help you making a gRPC server. Errors logged on [Sentry](https://sentry.io), Panic interceptor, HTTP status to gRPC code, ... Everything is made to assist you :)

Table of contents
=================

  * [Install](#install)
  * [Usage](#usage)
      * [Status](#status)
        * [Error](#error)
        * [New](#new)
        * [From error](#from-error)
        * [gRPC code from HTTP status](#grpc-code-from-http-status)
    * [Panic interceptor](#panic-interceptor)
  * [References](#references)

## Install

Use `go get` to install this package.

    go get github.com/visiperf/visigrpc


## Usage

### Status

#### Error

The `Error(code codes.Code, msg string) error` function is used to return a gRPC error and log it into [Sentry](https://sentry.io).

```go
type server struct { }

func main() {
  // init Sentry config
  if err := sentry.Init(sentry.ClientOptions{
    Dsn:         SENTRY_DSN,
    Environment: SENTRY_ENV,
    Transport:   sentry.NewHTTPSyncTransport(),
  }); err != nil {
    ...
  }
  
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
  return nil, status.Error(codes.Unimplemented, "implement me")
}
```

##### IMPORTANT : Only `Unknown`, `Internal` and `DataLoss` errors will be reported in Sentry !

#### New

The `New(code codes.Code, msg string) *spb.Status` function has same process as `Error(...) error` function, but returns a `*spb.Status` instance instead of `error`.

```go
type server struct { }

func main() {
  // init Sentry config
  if err := sentry.Init(sentry.ClientOptions{
    Dsn:         SENTRY_DSN,
    Environment: SENTRY_ENV,
    Transport:   sentry.NewHTTPSyncTransport(),
  }); err != nil {
    ...
  }
  
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
  st := status.New(codes.Internal, "Oups, an error !")
  // st.Code -> codes.Internal
  // st.Message -> "Oups, an error !"
  ...
  
  return nil, status.Error(codes.Unimplemented, "implement me")
}
```



#### From Error

If you receive a gRPC error (made with status.Error(...)), you can decode it with `FromError(err error) *spb.Status` to retrieve the gRPC code and message.

```go
type server struct { }

func main() {
  // init Sentry config
  if err := sentry.Init(sentry.ClientOptions{
    Dsn:         SENTRY_DSN,
    Environment: SENTRY_ENV,
    Transport:   sentry.NewHTTPSyncTransport(),
  }); err != nil {
    ...
  }
  
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
  return nil, status.Error(codes.Unimplemented, "implement me")
}

func (s *server) Create(ctx context.Context, in *CreateRequest) (*CreateResponse, error) {
  resp, err := s.Get(ctx, &GetRequest{}) // just for example, never directly call self functions with `s *server` !
  if err != nil {
    st := status.FromError(err)
    // st.Code -> codes.Unimplemented
    // st.Message -> "implement me"
    ...
  }
  
  return nil, status.Error(codes.Unimplemented, "implement me")
}
```

#### gRPC code from HTTP status

If you make an HTTP request, you can use the `GrpcCodeFromHttpStatus(status int) codes.Code` func to convert HTTP status code in response to gRPC code.

```go
code := status.GrpcCodeFromHttpStatus(http.StatusForbidden) // http status -> 403 (Forbidden)

// code -> 7 (codes.PermissionDenied)
```

### Panic interceptor

Visigrpc also provide a `RecoveryHandler` to capture and log panics for `Unary` and `Stream` functions.

```go
type server struct { }

func main() {
  // init Sentry config
  if err := sentry.Init(sentry.ClientOptions{
    Dsn:         SENTRY_DSN,
    Environment: SENTRY_ENV,
    Transport:   sentry.NewHTTPSyncTransport(),
  }); err != nil {
    ...
  }
  
  // gRPC server
  lis, err := net.Listen("tcp", ":9090")
  if err != nil {
    ...
  }
  
  var opts = []grpc_recovery.Option{
    grpc_recovery.WithRecoveryHandler(visigrpc.RecoveryHandler),
  }

  s := grpc.NewServer(
    grpc_middleware.WithUnaryServerChain(
       grpc_recovery.UnaryServerInterceptor(opts...),
    ),
    grpc_middleware.WithStreamServerChain(
       grpc_recovery.StreamServerInterceptor(opts...),
    ),
  )
  
  RegisterServiceServer(s, &server{})
  
  if err := s.Serve(lis); err != nil {
    ...
  }
}

func (s *server) Get(ctx context.Context, in *GetRequest) (*GetResponse, error) {
  panic("implement me") // will return codes.Unknown with message "implement me" and log error on Sentry
}
```



## References

* Sentry : [github.com/getsentry/sentry-go](https://github.com/getsentry/sentry-go)
* Go gRPC Middleware : [github.com/grpc-ecosystem/go-grpc-middleware](https://github.com/grpc-ecosystem/go-grpc-middleware)

