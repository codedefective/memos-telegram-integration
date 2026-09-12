package main

import (
	"context"
	"log/slog"
	"time"

	"github.com/usememos/memogram"
)

const (
	maxStartAttempts   = 36
	startRetryInterval = 5 * time.Second
)

func main() {
	ctx := context.Background()

	var service *memogram.Service
	for attempt := 1; attempt <= maxStartAttempts; attempt++ {
		var err error
		service, err = memogram.NewService()
		if err == nil {
			break
		}
		slog.Warn("failed to create service, retrying...",
			slog.Int("attempt", attempt),
			slog.String("error", err.Error()),
		)
		time.Sleep(startRetryInterval)
	}
	if service == nil {
		slog.Error("failed to create service after retries")
		return
	}

	service.Start(ctx)
}
