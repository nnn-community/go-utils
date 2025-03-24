package fetch

import (
    "encoding/json"
    "fmt"
    "net/url"
    "reflect"
)

func MakeQueryUrl(urlStr string, payload *interface{}) string {
    if payload == nil {
        return urlStr
    }

    parsedURL, err := url.Parse(urlStr)

    if err != nil {
        return urlStr
    }

    queryParams := url.Values{}
    v := reflect.ValueOf(payload)

    if v.Kind() == reflect.Ptr {
        v = v.Elem()
    }

    if v.Kind() != reflect.Struct {
        jsonBytes, err := json.Marshal(payload)

        if err != nil {
            return urlStr
        }

        var payloadMap map[string]interface{}
        err = json.Unmarshal(jsonBytes, &payloadMap)

        if err != nil {
            return urlStr
        }

        for key, value := range payloadMap {
            queryParams.Add(key, fmt.Sprintf("%v", value))
        }
    } else {
        t := v.Type()

        for i := 0; i < v.NumField(); i++ {
            field := v.Field(i)
            fieldName := t.Field(i).Name
            queryParams.Add(fieldName, fmt.Sprintf("%v", field.Interface()))
        }
    }

    parsedURL.RawQuery = queryParams.Encode()

    return parsedURL.String()
}
