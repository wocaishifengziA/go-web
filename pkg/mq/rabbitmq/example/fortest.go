package rabbitmq

import (
	"log"
	"time"

	"github.com/streadway/amqp"
)

func DoMq() {
	addr := "amqp://root:root@192.168.2.125:5672/"
	conn, err := amqp.Dial(addr)
	if err != nil {
		log.Println("conn failed, ", err.Error())
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Println("ch failed, ", err.Error())
	}
	defer ch.Close()
	ex1 := "xy_ex_1"
	ex2 := "xy_ex_2"

	err = ch.ExchangeDeclare(
		ex1,      // name
		"topic", // type
		true,     // durable
		false,    // auto-deleted
		false,    // internal
		false,    // no-wait
		nil,      // arguments
	)
	if err != nil {
		log.Println("ex1 failed, ", err.Error())
	}
	err = ch.ExchangeDeclare(
		ex2,      // name
		"topic", // type
		true,     // durable
		false,    // auto-deleted
		false,    // internal
		false,    // no-wait
		nil,      // arguments
	)
	if err != nil {
		log.Println("ex2 failed, ", err.Error())
	}
	q1 := "xy_q_1"
	q2 := "xy_q_2"
	qa, err := ch.QueueDeclare(
		q1,    // name
		true,  // durable
		false, // delete when usused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		log.Println("q1 failed, ", err.Error())
	}
	qb, err := ch.QueueDeclare(
		q2,    // name
		true,  // durable
		false, // delete when usused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		log.Println("q2 failed, ", err.Error())
	}

	err = ch.QueueBind(
		qa.Name,  // queue name
		"test.xxx.*", // routing key
		ex1,      // exchange
		false,
		nil)
	if err != nil {
		log.Println("bind1 failed, ", err.Error())
	}

	err = ch.QueueBind(
		qb.Name,  // queue name
		"test.*", // routing key
		ex1,      // exchange
		false,
		nil)
	if err != nil {
		log.Println("bind2 failed, ", err.Error())
	}

	err = ch.Publish(
		ex1,    // exchange
		"test.abb", // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte("gogogo"),
		})
	if err != nil {
		log.Println("pub failed, ", err.Error())
	}
	for {
		time.Sleep(time.Duration(3) * time.Second)
		log.Println("run...")
	}
}
