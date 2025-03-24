package strings

import "strconv"

func ToInt(text string, fallbackValue int) int {
    i, err := strconv.Atoi(text)

    if err != nil {
        return fallbackValue
    } else {
        return i
    }
}

func ToFloat(text string, fallbackValue float64) float64 {
    f, err := strconv.ParseFloat(text, 64)

    if err != nil {
        return fallbackValue
    } else {
        return f
    }
}
