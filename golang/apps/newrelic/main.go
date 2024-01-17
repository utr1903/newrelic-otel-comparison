package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/utr1903/newrelic-otel-comparison/golang/apps/newrelic/client"
	"github.com/utr1903/newrelic-otel-comparison/golang/apps/newrelic/server"
)

func main() {
	app, err := newrelic.NewApplication(
		newrelic.ConfigFromEnvironment(),
		newrelic.ConfigDistributedTracerEnabled(true),
		newrelic.ConfigAppLogForwardingEnabled(false),
		newrelic.ConfigCodeLevelMetricsEnabled(false),
	)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Wait for successfull handshake
	err = app.WaitForConnection(5 * time.Second)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Instantiate & run client
	c := client.New(app)
	go c.Simulate()

	// Instantiate server
	s := server.New()
	http.HandleFunc(newrelic.WrapHandleFunc(app, "/api", s.Handler))

	// Start server
	err = http.ListenAndServe(":8081", nil)
	if err != nil {
		fmt.Println("Error starting the server:", err)
	}
}
