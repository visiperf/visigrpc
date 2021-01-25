package main

import (
	"log"

	"github.com/visiperf/visigrpc/v3/status"
	"google.golang.org/grpc/codes"

	"github.com/getsentry/sentry-go"
)

const sentryDsn = "YOUR_SENTRY_DSN"
const sentryEnv = "YOUR_SENTRY_ENV"

func main() {
	if err := sentry.Init(sentry.ClientOptions{
		Dsn:         sentryDsn,
		Environment: sentryEnv,
		Transport:   sentry.NewHTTPSyncTransport(),
	}); err != nil {
		log.Fatalf("failed to init sentry: %v", err)
	}

	status.Error(codes.Internal, "a sentry logged error")
}
