package fetch

import (
    "context"
    "fmt"
    "time"
)

var defaultRepeat = false
var defaultRepeatErrorCodes = []int{429, 500, 503}
var defaultRepeatTimeout = time.Minute
var defaultRepeatMaxBounces = 5
var defaultTimeout = 5 * time.Minute

type Option func(f *FetchClient)

func WithContext(ctx context.Context) Option {
    return func(f *FetchClient) {
        f.ctx = ctx
    }
}

func WithHeaders(h map[string]string) Option {
    return func(f *FetchClient) {
        for k, v := range h {
            f.headers[k] = v
        }
    }
}

func WithContentType(c string) Option {
    return func(f *FetchClient) {
        f.headers["Content-Type"] = c
    }
}

func WithBearerToken(t string) Option {
    return func(f *FetchClient) {
        f.headers["Authorization"] = fmt.Sprintf("Bearer %s", t)
    }
}

func WithCredentials() Option {
    return func(f *FetchClient) {
        f.headers["Credentials"] = "include"
    }
}

func WithRepeat() Option {
    return func(f *FetchClient) {
        f.repeat = true
    }
}

func WithRepeatErrorCode(errorCodes []int) Option {
    return func(f *FetchClient) {
        f.repeatErrorCodes = errorCodes
    }
}

func WithRepeatTimeout(timeout time.Duration) Option {
    return func(f *FetchClient) {
        f.repeatTimeout = timeout
    }
}

func WithRepeatMaxBounces(maxBounces int) Option {
    return func(f *FetchClient) {
        f.repeatMaxBounces = maxBounces
    }
}

func WithTimeout(timeout time.Duration) Option {
    return func(f *FetchClient) {
        f.timeout = timeout
    }
}

func applyOptions(opts ...Option) *FetchClient {
    f := &FetchClient{
        headers:          make(map[string]string),
        repeat:           defaultRepeat,
        repeatErrorCodes: defaultRepeatErrorCodes,
        repeatTimeout:    defaultRepeatTimeout,
        repeatMaxBounces: defaultRepeatMaxBounces,
        timeout:          defaultTimeout,
    }

    for _, opt := range opts {
        opt(f)
    }

    if f.ctx == nil {
        f.ctx = context.Background()
    }

    return f
}
