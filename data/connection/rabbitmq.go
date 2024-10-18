package connection

import (
	"context"
	"fmt"
	"ncobase/common/config"
	"ncobase/common/log"

	amqp "github.com/rabbitmq/amqp091-go"
)

// newRabbitMQConnection creates a new RabbitMQ connection
func newRabbitMQConnection(conf *config.RabbitMQ) (*amqp.Connection, error) {
	if conf == nil || conf.URL == "" {
		log.Infof(context.Background(), "RabbitMQ configuration is nil or empty")
		return nil, nil
	}

	url := fmt.Sprintf("amqp://%s:%s@%s/%s", conf.Username, conf.Password, conf.URL, conf.Vhost)
	conn, err := amqp.DialConfig(url, amqp.Config{
		Heartbeat: conf.HeartbeatInterval,
		Vhost:     conf.Vhost,
	})
	if err != nil {
		log.Errorf(context.Background(), "Failed to connect to RabbitMQ: %v", err)
		return nil, err
	}

	log.Infof(context.Background(), "Connected to RabbitMQ")
	return conn, nil
}
