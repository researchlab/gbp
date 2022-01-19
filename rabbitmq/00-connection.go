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

/*
.env.json
{
	"DialUrl":"amqp://user:pwd@ip:5672/"
}
*/
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
	channel, err := conn.Channel()
	if err != nil {
		log.Fatal("channel:", err)
	}
	channel.ExchangeDeclare("test", amqp.ExchangeTopic, true, false, false, false, nil)
	channel.QueueDeclare("testQueue", true, false, false, false, nil)
	channel.QueueBind("testQueue", "testKey", "test", false, nil)
	channel.Close()
	done := make(chan bool, 1)
	for {
		select {
		case <-done:
			fmt.Println("done")
			fmt.Println(time.Now())
			return
		default:
			go testNotCloseConn(conn, done)
			//fmt.Println("test")
			//go testCloseConn(done)
		}
		time.Sleep(time.Nanosecond)
	}
}

func testNotCloseConn(conn *amqp.Connection, done chan bool) {
	mutex.Lock()
	channel, err := conn.Channel()
	count++
	fmt.Println(count)
	if err != nil {
		fmt.Println("*****", count, "****")
		fmt.Println("channel:", err)
		done <- true
		return
	}
	mutex.Unlock()
	amqpTest := amqp.Publishing{
		ContentType: "application/json",
		Body:        []byte("{name:'china',host:'localhost'}"),
	}

	err = channel.Publish("test", "testKey", false, false, amqpTest)
	if err != nil {
		fmt.Println("publish:", err)
		return
	}
	defer channel.Close()
	fmt.Println("发送成功！")
	time.Sleep(time.Second * 600)
}

func testCloseConn(done chan bool) {
	mutex.Lock()
	conn, err := amqp.Dial("amqp://admin:cl0udsuit1@172.118.59.90:5672/")
	count++
	if err != nil {
		fmt.Println("*****", count, "****")
		fmt.Println("dial:", err)
		done <- true
		return
	}
	mutex.Unlock() //加锁位置
	channel, err := conn.Channel()
	defer conn.Close()
	if err != nil {
		log.Println("channel:", err)
		return
	}
	defer channel.Close()

	amqpTest := amqp.Publishing{
		ContentType: "application/json",
		Body:        []byte("{name:'china',host:'localhost'}"),
	}
	channel.Qos(1, 0, true)
	err = channel.Publish("test", "testKey", false, false, amqpTest)
	if err != nil {
		log.Println("publish:", err)
	}
	fmt.Println(count, "  发送成功")
	time.Sleep(time.Second * 600) //睡眠协程，不让conn和channel关闭
}
