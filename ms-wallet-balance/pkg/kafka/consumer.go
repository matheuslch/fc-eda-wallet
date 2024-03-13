package kafka

import ckafka "github.com/confluentinc/confluent-kafka-go/kafka"

type Consumer struct {
	ConfigMap *ckafka.ConfigMap
	Topics    []string
	Consumer  *ckafka.Consumer
}

func NewKafkaConsumer(configMap *ckafka.ConfigMap, topics []string) (*Consumer, error) {
	c := &Consumer{
		ConfigMap: configMap,
		Topics:    topics,
	}

	var err error
	c.Consumer, err = ckafka.NewConsumer(c.ConfigMap)
	if err != nil {
		return nil, err
	}

	err = c.Consumer.SubscribeTopics(c.Topics, nil)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func (c *Consumer) Consume(msgChan chan *ckafka.Message) error {
	for {
		msg, err := c.Consumer.ReadMessage(-1)
		if err == nil {
			msgChan <- msg
		}
	}
}

func (c *Consumer) Close() {
	c.Consumer.Close()
}
