#!/bin/bash
migrate -path /app/internal/database/migrations -database "mysql://root:root@tcp(wallet-core-mysql:3307)/wallet?charset=utf8&parseTime=True&loc=Local" up \
& chmod +x /app/cmd/walletcore/main \
& ./cmd/walletcore/main \
& tail -f /dev/null
