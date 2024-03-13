package main

import (
	"database/sql"
	"fmt"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	_ "github.com/go-sql-driver/mysql"
	"github.com/matheuslch/fc-ms-wallet-balance/internal/consumer/handler"
	"github.com/matheuslch/fc-ms-wallet-balance/internal/database"
	"github.com/matheuslch/fc-ms-wallet-balance/internal/usecases/create_balance"
	"github.com/matheuslch/fc-ms-wallet-balance/internal/usecases/find_balance"
	"github.com/matheuslch/fc-ms-wallet-balance/internal/web"
	"github.com/matheuslch/fc-ms-wallet-balance/internal/web/webserver"
	"github.com/matheuslch/fc-ms-wallet-balance/pkg/kafka"
)

func main() {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", "root", "root", "wallet-balance-mysql", "3308", "wallet"))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	balanceDb := database.NewBalanceDB(db)
	findBalanceUseCase := find_balance.NewFindBalanceUseCase(balanceDb)
	createBalanceUseCase := create_balance.NewCreateBalanceUseCase(balanceDb)

	configMap := &ckafka.ConfigMap{
		"bootstrap.servers": "kafka:29092",
		"group.id":          "wallet",
	}
	cons, _ := kafka.NewKafkaConsumer(configMap, []string{"balances"})

	msgChan := make(chan *ckafka.Message)
	createBalanceKafkaHandler := handler.NewCreateBalanceKafkaHandler(msgChan)

	go cons.Consume(msgChan)
	go createBalanceKafkaHandler.Handle(createBalanceUseCase)

	webserver := webserver.NewWebServer(":3003")

	balanceHandler := web.NewWebBalanceHandler(*findBalanceUseCase)

	webserver.AddHandler("/balances/{account_id}", "get", balanceHandler.GetBalance)
	fmt.Println("Server is running")
	webserver.Start()
}
