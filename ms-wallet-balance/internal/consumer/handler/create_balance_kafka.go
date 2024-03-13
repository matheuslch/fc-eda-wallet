package handler

import (
	"encoding/json"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/matheuslch/fc-ms-wallet-balance/internal/usecases/create_balance"
)

type CreateBalanceKafkaHandler struct {
	MsgChan chan *ckafka.Message
}

func NewCreateBalanceKafkaHandler(msgChan chan *ckafka.Message) *CreateBalanceKafkaHandler {
	return &CreateBalanceKafkaHandler{
		MsgChan: msgChan,
	}
}

func (h *CreateBalanceKafkaHandler) Handle(createBalanceUseCase *create_balance.CreateBalanceUseCase) {
	for msg := range h.MsgChan {
		var dto create_balance.ConsumerDTO
		err := json.Unmarshal(msg.Value, &dto)
		if err == nil {
			createBalanceUseCase.Execute(dto.Payload)
		}
	}
}
