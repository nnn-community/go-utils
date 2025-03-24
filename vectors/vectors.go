package vectors

import (
    "fmt"
    "strconv"
    "strings"
)

func ToVector(embedding []float32) string {
    return fmt.Sprintf("[%s]", strings.Trim(strings.Join(strings.Fields(fmt.Sprint(embedding)), ","), "[]"))
}

func ToEmbedding(value string) []float32 {
    vectorString := strings.Trim(value, "[]")
    vectorValues := strings.Split(vectorString, ",")
    var embedding []float32

    for _, v := range vectorValues {
        f, err := strconv.ParseFloat(strings.TrimSpace(v), 32)

        if err != nil {
            return nil
        }

        embedding = append(embedding, float32(f))
    }

    return embedding
}
