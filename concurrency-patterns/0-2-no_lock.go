package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("生成3张票，卖出2张，至少剩余一张")
	wg := &sync.WaitGroup{}
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			Push(i)
		}(i)
	}
	for i := 0; i < 2; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			Pop(i)
		}(i)
	}
	wg.Wait()
	fmt.Println("剩余:", Total(), " 票:", Tickets())

	Reset()

	fmt.Println("生成2张票，卖出3张，至少卖出失败一张")
	for i := 0; i < 2; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			Push(i)
		}(i)
	}
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			Pop(i)
		}(i)
	}
	wg.Wait()
	fmt.Println("剩余:", Total(), " 票:", Tickets())
}

var total int64
var tickets []ticket
var pipe chan payload

// ticket 票模型
type ticket struct {
	id  int // ticket id
	opt int // opt type   1: push 生成票, 2:pop 买票
}

// payload 生成票模型
type payload struct {
	in  ticket
	out chan interface{}
}

// popload 购买票模型
type popload struct {
	ok  bool
	out ticket
}

func init() {
	pipe = make(chan payload, 10)
	go ticketer()
}

// Push 生成票
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

// Pop 购买票
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

// ticketer  串行进票出票
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

// --- 其它 functions

func Reset() {
	total = 0
	tickets = []ticket{}
}

func Total() int64 {
	return total
}

func Tickets() []ticket {
	return tickets
}
