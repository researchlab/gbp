package shape

import "math"

// interface 是用在调用的时候，收敛不同的类的
// 不同的类都有同一个方法，那么通过interface 定义一个类型， 在调用的时候就用这个类型，可以不用管是什么类型都可以调用这个方法
type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}
