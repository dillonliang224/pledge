package rabbitmq

import "log"

func Receive()  {
	conn, err := Connection()
	failOnError(err, "conn has error")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "chan has error")
	defer ch.Close()

	q, err := ch.QueueDeclare("hello", false, false, false, false, nil)
	failOnError(err, "declare queue error")

	msgs, err := ch.Consume(q.Name, "", false, false, false, false, nil)
	failOnError(err, "fail to registry consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf(" %s", d.Body)
		}
	}()

	<-forever
}

