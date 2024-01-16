package client

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/newrelic/go-agent/v3/newrelic"
)

type Client struct {
	app    *newrelic.Application
	client *http.Client
}

func New(app *newrelic.Application) *Client {
	return &Client{
		app:    app,
		client: &http.Client{},
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

	// Make request every second
	time.Sleep(time.Duration(time.Second))

	// Start parent transaction
	txn := c.app.StartTransaction("client")
	defer txn.End()

	// Prepare external request
	req, err := http.NewRequest("GET", "http://localhost:8080/api", nil)
	if err != nil {
		// Record error
		txn.NoticeError(err)
		fmt.Println(err)
		return
	}

	// Start external segment
	seg := newrelic.StartExternalSegment(txn, req)
	defer seg.End()

	// Perform HTTP request
	res, err := c.client.Do(req)
	if err != nil {
		// Record error
		txn.NoticeError(err)
		fmt.Println(err)
		return
	}
	io.ReadAll(res.Body)
	res.Body.Close()

	fmt.Println("Request succeeded.")
}
