package main

import "log/slog"

func main() {
	slog.Info("starting")
	slog.Info("done")
}

func add(a, b int) int {
	return a + b + 2
}
