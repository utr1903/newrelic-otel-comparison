package client

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
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

	// // Create span
	// ctx, span := otel.Tracer("client").Start(ctx, "say hello")
	// defer span.End()

	// Prepare request with span context
	req, _ := http.NewRequestWithContext(ctx, "GET", "http://localhost:8081/api", nil)

	// Perform HTTP request
	res, err := c.client.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	io.ReadAll(res.Body)
	res.Body.Close()

	fmt.Println("Request succeeded.")
}
