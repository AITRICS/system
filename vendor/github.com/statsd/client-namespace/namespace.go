package namespace

import "github.com/statsd/client-interface"
import "time"

type Client struct {
	c      statsd.Client
	prefix string
}

func New(client statsd.Client, prefix string) *Client {
	return &Client{client, prefix}
}

func (c *Client) Gauge(name string, value int) error {
	return c.c.Gauge(c.prefix+"."+name, value)
}

func (c *Client) Incr(name string) error {
	return c.c.Incr(c.prefix + name)
}

func (c *Client) IncrBy(name string, value int) error {
	return c.c.IncrBy(c.prefix+"."+name, value)
}

func (c *Client) Decr(name string) error {
	return c.c.Decr(c.prefix + name)
}

func (c *Client) DecrBy(name string, value int) error {
	return c.c.DecrBy(c.prefix+"."+name, value)
}

func (c *Client) Duration(name string, duration time.Duration) error {
	return c.c.Duration(c.prefix+"."+name, duration)
}

func (c *Client) Histogram(name string, value int) error {
	return c.c.Histogram(c.prefix+"."+name, value)
}

func (c *Client) Annotate(name string, value string, args ...interface{}) error {
	return c.c.Annotate(name, value, args...)
}

func (c *Client) Flush() error {
	return c.c.Flush()
}
