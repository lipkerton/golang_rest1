package main
import (
    "context"
    "errors"
    "log/slog"
    "net/http"
    "os"
    "os/signal"
    "syscall"

    "github.com/lipkerton/wildcard/internal/app"
    "github.com/lipkerton/wildcard/internal/config"
)

func main() {
    cfg := config.load()
}
