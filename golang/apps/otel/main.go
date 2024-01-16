package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/utr1903/newrelic-otel-comparison/golang/apps/otel/otel"
	"github.com/utr1903/newrelic-otel-comparison/golang/apps/otel/server"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/contrib/instrumentation/runtime"
)

func main() {

	ctx := context.Background()

	// Create tracer provider
	tp := otel.NewTraceProvider(ctx)
	defer otel.ShutdownTraceProvider(ctx, tp)

	// Create metric provider
	mp := otel.NewMetricProvider(ctx)
	defer otel.ShutdownMetricProvider(ctx, mp)

	// Start runtime metric collection
	err := runtime.Start(runtime.WithMinimumReadMemStatsInterval(time.Second))
	if err != nil {
		panic(err)
	}

	// Instantiate server
	srv := server.New()
	http.Handle("/api", otelhttp.NewHandler(http.HandlerFunc(srv.Handler), "api"))

	// Start server
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting the server:", err)
	}
}
