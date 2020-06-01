package lib

import "fmt"

var total int64
var tickets []ticket

func Reset() {
	total = 0
	tickets = []ticket{}
}

type ticket struct {
	id  int // ticket id
	opt int // opt type   1: push 生成票, 2:pop 买票
}

var pipe chan payload

type payload struct {
	in  ticket
	out chan interface{}
}

type popload struct {
	ok  bool
	out ticket
}

func init() {
	pipe = make(chan payload, 10)
	go ticketer()
}
func Push(i int) {
	// 生成接收者
	notify := make(chan interface{})
	// 生成票
	pipe <- payload{in: ticket{id: i, opt: 1}, out: notify}
	// 结果接收者
	n := <-notify
	if _, ok := n.(bool); ok {
		fmt.Println("Pusher:", i, " 生成票:", i, " 成功")
	} else {
		fmt.Println("Pusher:", i, " 生成票:", i, " 失败")
	}
}

func Pop(i int) {
	// 生成接收者
	notify := make(chan interface{})

	//购买票
	pipe <- payload{in: ticket{id: i, opt: 2}, out: notify}

	//结果接收者
	n := <-notify
	if out, ok := n.(popload); ok && out.ok {
		fmt.Println("Poper:", i, " 购票:", out.out, " 成功")
	} else {
		fmt.Println("Poper:", i, " 购票: 失败")
	}
}

func Total() int64 {
	return total
}

func Tickets() []ticket {
	return tickets
}

func ticketer() {
	for {
		select {
		case pl := <-pipe:
			if pl.in.opt == 1 { // 生成票
				tickets = append(tickets, pl.in)
				total++
				pl.out <- true
			} else if pl.in.opt == 2 { // 购票
				if total <= 0 {
					pl.out <- popload{ok: false}
				} else {
					pl.out <- popload{ok: true, out: tickets[0]}
					total--
					if total > 0 {
						tickets = tickets[1:]
					} else {
						tickets = []ticket{}
					}
				}
			}
		}
	}
}
