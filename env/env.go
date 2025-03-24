package env

import (
    "fmt"
    "log"
    "os"

    "github.com/joho/godotenv"
)

func Load() {
    env := os.Getenv("APP_ENV")

    if env == "" {
        env = "development"
    }

    if err := godotenv.Load(); err != nil {
        log.Fatal("Error loading .env file")
    }

    if err := godotenv.Load(fmt.Sprintf(".env.%s", env)); err != nil {
        log.Fatal(fmt.Sprintf("Error loading .env.%s file", env))
    }

    if _, err := os.Stat(".env.local"); err == nil {
        if err := godotenv.Load(".env.local"); err != nil {
            log.Fatal("Error loading .env.local file")
        }
    }
}

func LoadSingle() {
    if err := godotenv.Load(); err != nil {
        log.Fatal("Error loading .env file")
    }

    if _, err := os.Stat(".env.local"); err == nil {
        if err := godotenv.Load(".env.local"); err != nil {
            log.Fatal("Error loading .env.local file")
        }
    }
}

func IsDev() bool {
    env := os.Getenv("APP_ENV")

    return env == "" || env == "development"
}

func IsLocal() bool {
    return IsDev()
}

func IsStaging() bool {
    env := os.Getenv("APP_ENV")

    return env == "staging"
}

func IsProduction() bool {
    env := os.Getenv("APP_ENV")

    return env == "production"
}
