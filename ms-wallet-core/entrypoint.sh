#!/bin/bash
migrate -path /app/internal/database/migrations -database "mysql://root:root@tcp(wallet-core-mysql:3307)/wallet?charset=utf8&parseTime=True&loc=Local" up \
& go mod tidy \
& go run /app/cmd/walletcore/main.go \
& tail -f /dev/null
