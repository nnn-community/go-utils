package fetch

import (
    "bytes"
    "encoding/json"
    "fmt"
    "io"
    "net/http"
)

type Fetch interface {
    SetHeader(key string, value string)
    SetHeaders(headers map[string]string)
    SetContentType(contentType string)
    SetBearer(token string)
    Do(result *interface{}) error
}

func New(method string, url string, data ...interface{}) Fetch {
    var d *interface{} = nil

    if len(data) > 0 {
        d = &data[0]
    }

    if method == http.MethodGet {
        url = MakeQueryUrl(url, d)
        d = nil
    }

    return &FetchClient{
        method:  method,
        url:     url,
        data:    d,
        headers: map[string]string{},
    }
}

func (f *FetchClient) SetHeader(key string, value string) {
    f.headers[key] = value
}

func (f *FetchClient) SetHeaders(headers map[string]string) {
    for key, value := range headers {
        f.SetHeader(key, value)
    }
}

func (f *FetchClient) SetContentType(contentType string) {
    f.SetHeader("Content-Type", contentType)
}

func (f *FetchClient) SetBearer(token string) {
    f.SetHeader("Authorization", fmt.Sprintf("Bearer %s", token))
}

func (f *FetchClient) Do(result *interface{}) error {
    var err error

    payloadBytes, err := json.Marshal(f.data)

    if err != nil {
        return err
    }

    req, err := http.NewRequest(f.method, f.url, bytes.NewBuffer(payloadBytes))

    if err != nil {
        return err
    }

    for key, value := range f.headers {
        req.Header.Set(key, value)
    }

    client := &http.Client{}
    resp, err := client.Do(req)

    if err != nil {
        return err
    }

    body, err := io.ReadAll(resp.Body)

    if err != nil {
        return err
    }

    return json.Unmarshal(body, &result)
}

type FetchClient struct {
    method  string
    url     string
    data    *interface{}
    headers map[string]string
}

var _ Fetch = &FetchClient{}
