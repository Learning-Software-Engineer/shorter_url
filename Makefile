server:
	go run cmd/main.go
wire:
	cd internal/wire && wire
.PHONY: server wire