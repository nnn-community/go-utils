package fetch

import (
    "bytes"
    "context"
    "encoding/json"
    "errors"
    "github.com/nnn-community/go-utils/arrays"
    "io/ioutil"
    "net/http"
    "time"
)

type Fetch interface {
    Get(url string, payload *interface{}) ([]byte, error)
    Post(url string, payload *interface{}) ([]byte, error)
    Head(url string, payload *interface{}) ([]byte, error)
    Put(url string, payload *interface{}) ([]byte, error)
    Patch(url string, payload *interface{}) ([]byte, error)
    Delete(url string, payload *interface{}) ([]byte, error)
    Connect(url string, payload *interface{}) ([]byte, error)
    Options(url string, payload *interface{}) ([]byte, error)
    Trace(url string, payload *interface{}) ([]byte, error)
}

func New(opts ...Option) Fetch {
    return applyOptions(opts...)
}

func (f *FetchClient) Get(url string, payload *interface{}) ([]byte, error) {
    queriedUrl := MakeQueryUrl(url, payload)

    f.repeatCurrentBounces = 1

    return f.doRequest(queriedUrl, http.MethodGet, nil)
}

func (f *FetchClient) Post(url string, payload *interface{}) ([]byte, error) {
    payloadBytes, err := json.Marshal(payload)

    if err != nil {
        return nil, err
    }

    f.repeatCurrentBounces = 1

    return f.doRequest(url, http.MethodPost, &payloadBytes)
}

func (f *FetchClient) Head(url string, payload *interface{}) ([]byte, error) {
    payloadBytes, err := json.Marshal(payload)

    if err != nil {
        return nil, err
    }

    f.repeatCurrentBounces = 1

    return f.doRequest(url, http.MethodHead, &payloadBytes)
}

func (f *FetchClient) Put(url string, payload *interface{}) ([]byte, error) {
    payloadBytes, err := json.Marshal(payload)

    if err != nil {
        return nil, err
    }

    f.repeatCurrentBounces = 1

    return f.doRequest(url, http.MethodPut, &payloadBytes)
}

func (f *FetchClient) Patch(url string, payload *interface{}) ([]byte, error) {
    payloadBytes, err := json.Marshal(payload)

    if err != nil {
        return nil, err
    }

    f.repeatCurrentBounces = 1

    return f.doRequest(url, http.MethodPatch, &payloadBytes)
}

func (f *FetchClient) Delete(url string, payload *interface{}) ([]byte, error) {
    payloadBytes, err := json.Marshal(payload)

    if err != nil {
        return nil, err
    }

    f.repeatCurrentBounces = 1

    return f.doRequest(url, http.MethodDelete, &payloadBytes)
}

func (f *FetchClient) Connect(url string, payload *interface{}) ([]byte, error) {
    payloadBytes, err := json.Marshal(payload)

    if err != nil {
        return nil, err
    }

    f.repeatCurrentBounces = 1

    return f.doRequest(url, http.MethodConnect, &payloadBytes)
}

func (f *FetchClient) Options(url string, payload *interface{}) ([]byte, error) {
    payloadBytes, err := json.Marshal(payload)

    if err != nil {
        return nil, err
    }

    f.repeatCurrentBounces = 1

    return f.doRequest(url, http.MethodOptions, &payloadBytes)
}

func (f *FetchClient) Trace(url string, payload *interface{}) ([]byte, error) {
    payloadBytes, err := json.Marshal(payload)

    if err != nil {
        return nil, err
    }

    f.repeatCurrentBounces = 1

    return f.doRequest(url, http.MethodTrace, &payloadBytes)
}

func (f *FetchClient) doRequest(url string, method string, body *[]byte) ([]byte, error) {
    var rb *bytes.Reader

    if body == nil {
        rb = bytes.NewReader(nil)
    } else {
        rb = bytes.NewReader(*body)
    }

    req, err := http.NewRequestWithContext(f.ctx, method, url, rb)

    if err != nil {
        return nil, err
    }

    for k, v := range f.headers {
        req.Header.Set(k, v)
    }

    client := &http.Client{Timeout: f.timeout}
    r, err := client.Do(req)

    if err != nil {
        return nil, err
    }

    defer r.Body.Close()

    b, err := ioutil.ReadAll(r.Body)

    if err != nil {
        return nil, err
    }

    if r.StatusCode != http.StatusOK {
        canRepeat := f.repeat && arrays.Contains(f.repeatErrorCodes, r.StatusCode) && f.repeatCurrentBounces <= f.repeatMaxBounces

        if canRepeat {
            f.repeatCurrentBounces += 1
            time.Sleep(f.repeatTimeout)

            return f.doRequest(url, method, body)
        }

        return nil, errors.New(r.Status)
    }

    f.repeatCurrentBounces = 1

    return b, nil
}

type FetchClient struct {
    ctx                  context.Context
    headers              map[string]string
    repeat               bool
    repeatErrorCodes     []int
    repeatTimeout        time.Duration
    repeatCurrentBounces int
    repeatMaxBounces     int
    timeout              time.Duration
}

var _ Fetch = &FetchClient{}
