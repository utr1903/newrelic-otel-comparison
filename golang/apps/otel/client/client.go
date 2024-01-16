package client

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel"
)

type Client struct {
	client *http.Client
}

func New() *Client {
	return &Client{
		client: &http.Client{Transport: otelhttp.NewTransport(http.DefaultTransport)},
	}
}

func (c *Client) Simulate() {

	go func() {
		for {
			c.request()
		}
	}()
}

func (c *Client) request() {

	fmt.Println("Requesting...")

	// Get context
	ctx := context.Background()

	// Make request every second
	time.Sleep(time.Duration(time.Second))

	// Create parent span
	ctx, span := otel.Tracer("client").Start(ctx, "call-to-newrelic-app")
	defer span.End()

	// Prepare request with span context -> creates client span
	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8081/api", nil)
	if err != nil {
		// Record error
		span.RecordError(err)
		fmt.Println(err)
		return
	}

	// Perform HTTP request
	res, err := c.client.Do(req)
	if err != nil {
		// Record error
		span.RecordError(err)
		fmt.Println(err)
		return
	}

	io.ReadAll(res.Body)
	res.Body.Close()

	fmt.Println("Request succeeded.")
}
