package adapters

import (
    "apiInvitation/src/users/domain/entities"
    "context"
    "encoding/json"
    "log"
    "time"

    amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQAdapter struct {
    connectionString string
}

func NewRabbitMQAdapter(connectionString string) *RabbitMQAdapter {
    return &RabbitMQAdapter{connectionString: connectionString}
}

func (r *RabbitMQAdapter) failOnError(err error, msg string) {
    if err != nil {
        log.Panicf("%s: %s", msg, err)
    }
}

func (r *RabbitMQAdapter) Publish(message *entities.User) error {
    conn, err := amqp.Dial(r.connectionString)
    r.failOnError(err, "Failed to connect to RabbitMQ")
    defer conn.Close()

    ch, err := conn.Channel()
    r.failOnError(err, "Failed to open a channel")
    defer ch.Close()

    err = ch.ExchangeDeclare(
        "logs",   // name
        "direct", // type
        true,     // durable
        false,    // auto-deleted
        false,    // internal
        false,    // no-wait
        nil,      // arguments
    )
    r.failOnError(err, "Failed to declare an exchange")

    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    body, err := json.Marshal(message)
    r.failOnError(err, "Failed to marshal JSON")

    err = ch.PublishWithContext(ctx,
        "logs", // exchange
        "",     // routing key
        false,  // mandatory
        false,  // immediate
        amqp.Publishing{
            ContentType: "application/json",
            Body:        body,
        })
    r.failOnError(err, "Failed to publish a message")
    return nil
}
