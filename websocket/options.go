package websocket

import (
    "errors"
    "fmt"
)

var ErrInvalidOptions = errors.New("invalid options")

type Option func(f *WSClient)

func WithURL(url string) Option {
    return func(c *WSClient) {
        c.url = url
    }
}

func WithHeaders(h map[string]string) Option {
    return func(f *WSClient) {
        for k, v := range h {
            f.headers[k] = v
        }
    }
}

func WithBearerToken(t string) Option {
    return func(f *WSClient) {
        f.headers["Authorization"] = fmt.Sprintf("Bearer %s", t)
    }
}

func WithCredentials() Option {
    return func(f *WSClient) {
        f.headers["Credentials"] = "include"
    }
}

func applyOptions(opts ...Option) (*WSClient, error) {
    c := &WSClient{}

    for _, opt := range opts {
        opt(c)
    }

    if c.url == "" {
        return nil, fmt.Errorf("%w: missing URL", ErrInvalidOptions)
    }

    return c, nil
}
