package rabbitmq

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/streadway/amqp"
)

type Broker struct {
	exchange     string
	exchangeType string
	queueSuffix  string
	url          string
	conn         *amqp.Connection
	connClose    chan *amqp.Error

	channels sync.Map
}

func NewBroker(url, exchange, exchangeType, queueSuffix string) (*Broker, error) {
	conn, err := amqp.Dial(url)
	if err != nil {
		return nil, err
	}

	broker := &Broker{
		exchange:     exchange,
		exchangeType: exchangeType,
		queueSuffix:  queueSuffix,
		url:          url,
		conn:         conn,
		connClose:    conn.NotifyClose(make(chan *amqp.Error)),
	}

	// 启动一个goroutine，监听关闭信号，收到通知后，关闭连接
	go func() {
		<-broker.connClose
		broker.conn = nil
	}()

	ticker := time.NewTicker(5 * time.Second)
	// 启动一个goroutine，定时心跳检测mq服务
	go func() {
		for range ticker.C {
			broker.checkConnection()
		}
	}()

	return broker, nil
}

type HandleFlag string

const (
	HandleSuccess HandleFlag = "success"
	HandleDrop    HandleFlag = "drop"
	HandleRequeue HandleFlag = "requeue"
)

type Job interface {
	Handle([]byte) HandleFlag
}

// 注册worker，准备consume
func (b *Broker) Register(key string, concurrency int, job Job) {
	for {
		log.Printf("job %s starting...", key)

	}
}

func (b *Broker) readyConsume(key string, concurrency int, job Job) (bool, error) {
	channel, err := b.getChannel(key)
	if err != nil {
		return true, err
	}

	queue, err := b.declare(channel, key)
	if err != nil {
		return true, err
	}

	// qos, 限流，控制消费数量
	if err := channel.Qos(10, 0, false); err != nil {
		return true, fmt.Errorf("channel qos error: %s", err)
	}

	// 消息通道
	deliveries, err := channel.Consume(queue.Name, "", false, false, false, false, nil)

	// 关闭信号
	channelClose := channel.NotifyClose(make(chan *amqp.Error))

	// 有缓冲通道，控制并发数
	pool := make(chan struct{}, concurrency)

	go func() {
		for i := 0; i < concurrency; i++ {
			pool <- struct{}{}
		}
	}()

	for {
		select {
		case err := <-channelClose:
			b.channels.Delete(key)
			return true, fmt.Errorf("channel close: %s", err)
		case d := <-deliveries:
			if concurrency > 0 {
				<-pool
			}

			go func() {
				var flag HandleFlag

				switch flag = job.Handle(d.Body); flag {
				case HandleSuccess:
					d.Ack(false)
				case HandleDrop:
					d.Nack(false, false)
				case HandleRequeue:
					d.Nack(false, false)
				default:
					d.Nack(false, false)
				}

				if concurrency > 0 {
					pool <- struct{}{}
				}
			}()
		}
	}

	return false, nil
}

func (b *Broker) checkConnection() {
	if b.conn != nil {
		return
	}

	conn, err := amqp.Dial(b.url)
	if err != nil {
		log.Printf("broker redial faild: %v", err)
		return
	}

	b.connClose = conn.NotifyClose(make(chan *amqp.Error))
	go func() {
		<-b.connClose
		b.conn = nil
	}()

	b.conn = conn
}

// 发送消息
func (b *Broker) CreateTask(key string, data interface{}) error {
	channel, err := b.CreateChannel()
	if err != nil {
		return err
	}
	defer channel.Close()

	var body []byte
	if d, ok := data.(string); ok {
		body = []byte(d)
	} else {
		body, err = json.Marshal(data)
		if err != nil {
			return nil
		}
	}

	if _, err := b.declare(channel, key); err != nil {
		return err
	}

	return channel.Publish(b.exchange, key, false, false, amqp.Publishing{
		Headers:      amqp.Table{},
		ContentType:  "",
		Body:         body,
		DeliveryMode: amqp.Persistent,
	})
}

func (b *Broker) CreateChannel() (*amqp.Channel, error) {
	if b.conn == nil {
		return nil, errors.New("no available connection when create channel")
	}

	channel, err := b.conn.Channel()
	if err != nil {
		return nil, fmt.Errorf("create channel error: %s", err)
	}

	if err = channel.Confirm(false); err != nil {
		return nil, fmt.Errorf("channel not into confirm mode: %s", err)
	}

	return channel, nil
}

func (b *Broker) getChannel(key string) (*amqp.Channel, error) {
	if value, ok := b.channels.Load(key); ok {
		// 判断从map里取的值是否是*amqp.Channel类型
		if c, ok := value.(*amqp.Channel); ok {
			return c, nil
		}
	}

	channel, err := b.CreateChannel()
	if err != nil {
		return nil, err
	}

	b.channels.Store(key, channel)
	return channel, nil
}

func (b *Broker) declare(channel *amqp.Channel, key string) (amqp.Queue, error) {
	if err := channel.ExchangeDeclare(b.exchange, b.exchangeType, true, false, false, false, nil); err != nil {
		return amqp.Queue{}, fmt.Errorf("exchange declare error: %s", err)
	}

	queue, err := channel.QueueDeclare(fmt.Sprintf("%s_%s", key, b.queueSuffix), true, false, false, false, nil)
	if err != nil {
		return queue, fmt.Errorf("queue declare error: %s", err)
	}
	if err = channel.QueueBind(queue.Name, key, b.exchange, false, nil); err != nil {
		return queue, fmt.Errorf("queue bind error: %s", err)
	}
	return queue, nil
}
