package main

import (
	"errors"
	"log/slog"

	"github.com/joaoprofile/glog"
)

func main() {
	glog.New("API")
	err := errors.New("error test")

	glog.Info("my message with error", slog.Any("error", err))
}
