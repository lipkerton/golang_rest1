package config

import (
    "os"
    "time"
)

type Config struct {
    HTTPAddr        string
    ShutdownTimeout time.Duration  // пауза перед выключением сервера.
}

func Load() Config {
    addr := os.Getenv("HTTP_ADDR")
    if addr == "" {
        addr = ":" + getEnv("PORT", "8080")
    }
    return Config(
        HTTPAddr:        addr,
        ShutdownTimeout: 10 * time.Second
    )
}
