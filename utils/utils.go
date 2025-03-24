package utils

func Or[T any](value *T, fallback T) T {
    if value != nil {
        return *value
    }

    return fallback
}
