package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/streadway/amqp"
)

type C struct {
	DialUrl string
}

var c = C{}

func init() {
	fs, err := os.ReadFile(".env.json")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(fs, &c)
	if err != nil {
		panic(err)
	}
}

var count int
var mutex sync.Mutex

func main() {
	fmt.Println(time.Now())
	conn, err := amqp.Dial(c.DialUrl)
	if err != nil {
		log.Fatal("dial:", err)
	}
	ch, err := conn.Channel()
	if err != nil {
		log.Fatal("channel:", err)
	}
	q, err := ch.QueueDeclare(
		"test-queue",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatal("QueueDeclare:", err)
	}
	//err = ch.ExchangeDeclare("resource-topic", "topic", true, false, false, false, nil)
	//if err != nil {
	//	log.Fatal("ExchangeDeclare:", err)
	//}
	err = ch.QueueBind(q.Name, "resource.#", "resource-topic", false, nil)
	if err != nil {
		log.Fatal("QueueBind:", err)
	}
	go func() {
		for {
			d, err := ch.Consume(q.Name, "", false, false, false, false, nil)
			if err != nil {
				fmt.Errorf("channel.consume failed, err:%v", err)
				time.Sleep(1 * time.Second)
				continue
			}
			fmt.Println("waiting for msg ...")
			for msg := range d {
				fmt.Printf("%s\n", string(msg.Body))
				msg.Ack(true)
			}
			time.Sleep(1 * time.Second)
		}
	}()
	chh := make(chan bool)
	chh <- true
}
