package pub

import (
	"errors"
	"fmt"
	"time"

	"github.com/streadway/amqp"

	rmqpub "github.com/evmartinelli/go-rifa-microservice/pkg/rabbitmq/rmq_pub"
)

// ErrConnectionClosed -.
var ErrConnectionClosed = errors.New("rmq_pub - Connection closed")

const (
	_defaultWaitTime = 5 * time.Second
	_defaultAttempts = 10
	_defaultTimeout  = 2 * time.Second
)

// Client -.
type Client struct {
	conn           *rmqpub.Connection
	serverExchange string
	error          chan error
	stop           chan struct{}
	timeout        time.Duration
}

// New -.
func New(url, clientExchange string, opts ...Option) (*Client, error) {
	cfg := rmqpub.Config{
		URL:      url,
		WaitTime: _defaultWaitTime,
		Attempts: _defaultAttempts,
	}

	c := &Client{
		conn:           rmqpub.New("test", cfg),
		serverExchange: "test",
		error:          make(chan error),
		stop:           make(chan struct{}),
		timeout:        _defaultTimeout,
	}

	// Custom options
	for _, opt := range opts {
		opt(c)
	}

	err := c.conn.AttemptConnect()
	if err != nil {
		return nil, fmt.Errorf("rmq_rpc client - NewClient - c.conn.AttemptConnect: %w", err)
	}

	return c, nil
}

func (c *Client) Publish(body string) error {
	err := c.conn.Channel.Publish(c.serverExchange, "", false, false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	if err != nil {
		return fmt.Errorf("c.Channel.Publish: %w", err)
	}

	return nil
}
