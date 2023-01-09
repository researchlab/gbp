package bucket

import (
	"sync"
	"time"
)

// 简单实现令牌桶算法
// 实际情况可以直接调用开源的rate库

type bucket struct {
	rate  int           // 每分钟频率(每分钟加入多少个令牌)
	token chan struct{} // 存放令牌的地方
	cap   int           // 容量
	mu    *sync.Mutex   //桶内的锁
	pause bool          // 暂停
	stop  bool          // 停止
}

func NewBucket(rate, cap int) *bucket {
	if cap < 1 {
		panic("limit bucket cap error")
	}
	return &bucket{
		token: make(chan struct{}, cap),
		rate:  rate,
		mu:    new(sync.Mutex),
		cap:   cap,
	}
}

func (b *bucket) Start() {
	go b.addToken()
}

func (b *bucket) addToken() {
	for {
		b.mu.Lock()
		if b.stop {
			close(b.token)
			b.mu.Unlock()
			return
		}
		if b.pause {
			b.mu.Unlock()
			time.Sleep(time.Second)
			continue
		}
		b.token <- struct{}{}
		d := time.Minute / time.Duration(b.rate)
		b.mu.Unlock()
		time.Sleep(d)
	}
}

func (b *bucket) GetToken() {
	<-b.token
}

func (b *bucket) Pause() {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.pause = true
}

func (b *bucket) Pause() {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.pause = true
}

func (b *bucket) Stop() {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.stop = true
}

func (b *bucket) Reset() {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.token = make(chan struct{}, b.cap)
}
