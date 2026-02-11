package arrays

import (
    "errors"
    "fmt"
    "reflect"
    "sort"
)

func Filter[T any](collection []T, callback func(key int, value T) bool) []T {
    result := make([]T, 0)

    for i, item := range collection {
        if callback(i, item) {
            result = append(result, item)
        }
    }

    return result
}

func Find[T any](collection []T, callback func(key int, value T) bool) (T, error) {
    for i, item := range collection {
        if callback(i, item) {
            return item, nil
        }
    }

    var none T

    return none, errors.New("item not found")
}

func FindKey[T any](collection []T, callback func(key int, value T) bool) int {
    for i, item := range collection {
        if callback(i, item) {
            return i
        }
    }

    return -1
}

func Flatten(items interface{}, key string) []string {
    sliceValue := reflect.ValueOf(items)
    result := make([]string, sliceValue.Len())

    for i := 0; i < sliceValue.Len(); i++ {
        itemValue := sliceValue.Index(i)
        fieldValue := itemValue.FieldByName(key)

        if fieldValue.IsValid() {
            result[i] = fmt.Sprint(fieldValue.Interface())
        }
    }

    return result
}

func Group[T any](data []T, key string) map[string][]T {
    itemMap := make(map[string][]T)
    sliceValue := reflect.ValueOf(data)

    for i, item := range data {
        itemValue := sliceValue.Index(i)
        fieldValue := itemValue.FieldByName(key)
        id := fmt.Sprint(fieldValue.Interface())

        itemMap[id] = append(itemMap[id], item)
    }

    return itemMap
}

func GroupKey[T any](data []T, key string) []string {
    idMap := make(map[string]bool)
    uniqueIds := make([]string, 0)
    sliceValue := reflect.ValueOf(data)

    for i := range data {
        itemValue := sliceValue.Index(i)
        fieldValue := itemValue.FieldByName(key)
        id := fmt.Sprint(fieldValue.Interface())

        if _, exists := idMap[id]; !exists {
            idMap[id] = true
            uniqueIds = append(uniqueIds, id)
        }
    }

    return uniqueIds
}

func IndexOf[T comparable](arr []T, item T) int {
    for i, a := range arr {
        if a == item {
            return i
        }
    }

    return -1
}

func Contains[T comparable](arr []T, item T) bool {
    return IndexOf(arr, item) != -1
}

func Map[T any, R any](collection []T, callback func(key int, value T) R) []R {
    result := make([]R, 0)

    for i, item := range collection {
        result = append(result, callback(i, item))
    }

    return result
}

func Reverse[T any](items []T) []T {
    for i, j := 0, len(items)-1; i < j; i, j = i+1, j-1 {
        items[i], items[j] = items[j], items[i]
    }

    return items
}

func Sort[T any](collection []T, key string) SortResult[T] {
    result := make([]T, len(collection))
    copy(result, collection)

    sort.Slice(result, func(i, j int) bool {
        val1 := reflect.ValueOf(result[i])
        val2 := reflect.ValueOf(result[j])

        if val1.Kind() == reflect.Ptr {
            val1 = val1.Elem()
        }

        if val2.Kind() == reflect.Ptr {
            val2 = val2.Elem()
        }

        field1 := val1.FieldByName(key).String()
        field2 := val2.FieldByName(key).String()

        return field1 < field2
    })

    return SortResult[T]{
        collection: result,
    }
}

type SortResult[T any] struct {
    collection []T
}

func (s SortResult[T]) Asc() []T {
    return s.collection
}

func (s SortResult[T]) Desc() []T {
    return Reverse(s.collection)
}
