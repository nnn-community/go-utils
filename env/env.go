package env

import (
    "fmt"
    "log"
    "os"
    "path/filepath"

    "github.com/joho/godotenv"
)

func Load() {
    wd, err := os.Getwd()

    if err != nil {
        log.Fatal("Error getting working directory")
    }

    env := os.Getenv("APP_ENV")

    if env == "" {
        env = "development"
    }

    if err := godotenv.Load(filepath.Join(wd, ".env")); err != nil {
        log.Fatal("Error loading .env file")
    }

    if err := godotenv.Load(filepath.Join(wd, fmt.Sprintf(".env.%s", env))); err != nil {
        log.Fatal(fmt.Sprintf("Error loading .env.%s file", env))
    }

    if _, err := os.Stat(filepath.Join(wd, ".env.local")); err == nil {
        if err := godotenv.Load(filepath.Join(wd, ".env.local")); err != nil {
            log.Fatal("Error loading .env.local file")
        }
    }
}

func LoadSingle() {
    wd, err := os.Getwd()

    if err != nil {
        log.Fatal("Error getting working directory")
    }

    if err := godotenv.Load(filepath.Join(wd, ".env")); err != nil {
        log.Fatal("Error loading .env file")
    }

    if _, err := os.Stat(filepath.Join(wd, ".env.local")); err == nil {
        if err := godotenv.Load(filepath.Join(wd, ".env.local")); err != nil {
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
