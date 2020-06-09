package main

import (
	"log"

	"github.com/visiperf/visigrpc/v3/status"
	"google.golang.org/grpc/codes"

	"github.com/getsentry/sentry-go"
)

const SENTRY_DSN = "YOUR_SENTRY_DSN"
const SENTRY_ENV = "YOUR_SENTRY_ENV"

func main() {
	if err := sentry.Init(sentry.ClientOptions{
		Dsn:         SENTRY_DSN,
		Environment: SENTRY_ENV,
		Transport:   sentry.NewHTTPSyncTransport(),
	}); err != nil {
		log.Fatalf("failed to init sentry: %v", err)
	}

	status.Error(codes.Internal, "a sentry logged error")
}
