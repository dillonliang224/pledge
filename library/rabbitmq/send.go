package rabbitmq

import "github.com/streadway/amqp"

// https://www.rabbitmq.com/tutorials/tutorial-one-go.html

func Send(msg string)  {
    conn, err := Connection()
    failOnError(err, "conn has error")
    defer conn.Close()

    ch, err := conn.Channel()
    failOnError(err, "chan has error")
    defer ch.Close()

    q, err := ch.QueueDeclare("hello", false, false, false, false, nil)
    failOnError(err, "declare queue error")

    err = ch.Publish("", q.Name, false, false, amqp.Publishing{
        ContentType: "text/plain",
        Body: []byte(msg),
    })

    failOnError(err, "publish message error")
}