package rabbitmq

import (
	"log"
	"time"

	"github.com/streadway/amqp"
)

func DoConfirm() {
	// This example acts as a bridge, shoveling all messages sent from the source
	// exchange "log" to destination exchange "log".

	// Confirming publishes can help from overproduction and ensure every message
	// is delivered.

	// Setup the source of the store and forward
	
	source, err := amqp.Dial("amqp://root:root@192.168.2.125:5672/")
	if err != nil {
		log.Fatalf("connection.open source: %s", err)
	}
	defer source.Close()

	chs, err := source.Channel()
	if err != nil {
		log.Fatalf("channel.open source: %s", err)
	}

	if err := chs.ExchangeDeclare("log", "topic", true, false, false, false, nil); err != nil {
		log.Fatalf("exchange.declare destination: %s", err)
	}

	if _, err := chs.QueueDeclare("remote-tee", true, false, false, false, nil); err != nil {
		log.Fatalf("queue.declare source: %s", err)
	}

	if err := chs.QueueBind("remote-tee", "#", "log", false, nil); err != nil {
		log.Fatalf("queue.bind source: %s", err)
	}

	shovel, err := chs.Consume("remote-tee", "shovel", false, false, false, false, nil)
	if err != nil {
		log.Fatalf("basic.consume source: %s", err)
	}

	// Setup the destination of the store and forward
	destination, err := amqp.Dial("amqp://root:root@192.168.2.125:5672/")
	if err != nil {
		log.Fatalf("connection.open destination: %s", err)
	}
	defer destination.Close()

	chd, err := destination.Channel()
	if err != nil {
		log.Fatalf("channel.open destination: %s", err)
	}

	if err := chd.ExchangeDeclare("log", "topic", true, false, false, false, nil); err != nil {
		log.Fatalf("exchange.declare destination: %s", err)
	}
	
	// Buffer of 1 for our single outstanding publishing
	// confirms := chd.NotifyPublish(make(chan amqp.Confirmation, 1))

	// if err := chd.Confirm(false); err != nil {
	// 	log.Fatalf("confirm.select destination: %s", err)
	// }
	// ConfirmPush(chd, confirms)
	// ConfirmPush(chd, confirms)
	// ConfirmPush(chd, confirms)
	// ConfirmPush(chd, confirms)
	// ConfirmPush(chd, confirms)
	// ConfirmPush(chd, confirms)
	// ConfirmPush(chd, confirms)
	// ConfirmPush(chd, confirms)
	// ConfirmPush(chd, confirms)
	// ConfirmPush(chd, confirms)
	// ConfirmPush(chd, confirms)

	// Now pump the messages, one by one, a smarter implementation
	// would batch the deliveries and use multiple ack/nacks
	// count := 5
	// var msg = amqp.Delivery{}
	// var ok bool
	for {
		// count --
		msg, ok := <-shovel
		if !ok {
			log.Fatalf("source channel closed, see the reconnect example for handling this")
		}
		// err = chd.Publish("log", "xxx", false, false, amqp.Publishing{
		// 	ContentType: "text/plain",
		// 	Body:        []byte("gogogo"),
		// })

		// if err != nil {
		// 	msg.Nack(false, false)
		// 	log.Fatalf("basic.publish destination: %+v", msg)
		// }

		// only ack the source delivery when the destination acks the publishing
		// if confirmed := <-confirms; confirmed.Ack {
		// 	msg.Ack(false)
		// } else {
		// msg.Nack(false, false)
		// }
		log.Printf("get from mq: %+v", msg)
		// for {
		// 	log.Println("run...")
		// 	time.Sleep(time.Duration(3) * time.Second)
		// }
		msg.Nack(false, true)
	}
	// msg.Ack(true)
}

func ConfirmPush(chd *amqp.Channel, confirms chan amqp.Confirmation) bool {
	var err error
	for {
		err = chd.Publish("log", "xxx", false, false, amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte("gogogo"),
		})
		if err != nil {
			log.Fatalf("publish: %+v", err)
			continue
		}
		select {
		case confirm := <-confirms:
			if confirm.Ack {
				log.Printf("Push confirmed! num: %d\n", confirm.DeliveryTag)
				return true
			}
		case <-time.After(time.Duration(3) * time.Second):
			log.Println("Push didn't confirm. Retrying...")
		}
	}
}
