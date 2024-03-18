#!/bin/bash
migrate -path /app/internal/database/migrations -database "mysql://root:root@tcp(wallet-balance-mysql:3308)/wallet?charset=utf8&parseTime=True&loc=Local" up \
& go mod tidy \
& go run /app/cmd/walletbalance/main.go \
& tail -f /dev/null
