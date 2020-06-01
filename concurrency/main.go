package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"

	"github.com/researchlab/gbp/concurrency/lib"
)

//var url = "http://researchlab.github.io"
var url = "http://www.baidu.com"

func main() {
	VisitSite(url)
	fmt.Println("======concurrency sync mutex")
	Seller()
	fmt.Println("======concurrency no lock")
	fmt.Println("生成3张票，卖出2张，至少剩余一张")
	wg := &sync.WaitGroup{}
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			lib.Push(i)
		}(i)
	}
	for i := 0; i < 2; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			lib.Pop(i)
		}(i)
	}
	wg.Wait()
	fmt.Println("剩余:", lib.Total(), " 票:", lib.Tickets())
	fmt.Println("生成2张票，卖出3张，至少卖出失败一张")
	lib.Reset()
	for i := 0; i < 2; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			lib.Push(i)
		}(i)
	}
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			lib.Pop(i)
		}(i)
	}
	wg.Wait()
	fmt.Println("剩余:", lib.Total(), " 票:", lib.Tickets())
}

//  VisitSite Http Request with select timeout
func VisitSite(url string) {
	res, err := lib.HttpGet(&lib.GetOptions{Url: url, Timeout: 10})
	if err != nil {
		return
	}
	siteData := new(interface{})
	err = res.Unmarshal(siteData)
	if err != nil {
		return
	}
}

// Seller  sell 5 tickets
func Seller() {

	// 设置真正意义上的并发
	runtime.GOMAXPROCS(runtime.NumCPU())

	wg := sync.WaitGroup{}

	wg.Add(5)
	// 生成随机种子
	rand.Seed(time.Now().Unix())

	// 并发5个goroutine来卖票
	for i := 0; i < 5; i++ {
		go lib.SellTickets(&wg, i)
	}

	wg.Wait()
	// 退出时打印还有多少票
	fmt.Println(lib.TotalTickets(), "done")
}
