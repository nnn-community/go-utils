package websocket

import (
    "encoding/json"
    "github.com/gorilla/websocket"
    "net/http"
)

type WS interface {
}

func New(opts ...Option) (WS, error) {
    c, err := applyOptions(opts...)

    if err != nil {
        return nil, err
    }

    _, err = c.reconnect()

    if err != nil {
        return nil, err
    }

    return c, nil
}

func (c WSClient) Send(message *interface{}) error {
    payloadBytes, err := json.Marshal(message)

    if err != nil {
        return err
    }

    client, err := c.reconnect()

    if err != nil {
        return err
    }

    return client.WriteMessage(websocket.TextMessage, payloadBytes)
}

func (c WSClient) reconnect() (*websocket.Conn, error) {
    if c.client != nil {
        if _, _, err := c.client.ReadMessage(); err == nil {
            return c.client, nil
        }

        c.client.Close()
    }

    headers := http.Header{}

    for k, v := range c.headers {
        headers.Set(k, v)
    }

    newClient, _, err := websocket.DefaultDialer.Dial(c.url, headers)

    if err != nil {
        return nil, err
    }

    c.client = newClient

    return c.client, nil
}

type WSClient struct {
    url     string
    headers map[string]string
    client  *websocket.Conn
}

var _ WS = &WSClient{}
