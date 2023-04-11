package test

import (
	"fmt"
	"testing"
)

type Queue struct {
	content []byte
	pos     int
}

func (q *Queue) ReadUnsafe(size int) []byte {
	if q.pos+size >= len(q.content) {
		return nil
	}
	pos := q.pos
	q.pos = q.pos + size
	return q.content[pos:q.pos]
}

func (q *Queue) ReadSafe(size int) []byte {
	if q.pos+size >= len(q.content) {
		return nil
	}
	pos := q.pos
	q.pos = q.pos + size

	// 注意copy 时，ret一定要声明了大小，否则copy 没有用
	ret := make([]byte, size)
	copy(ret, q.content[pos:q.pos])
	return ret
}

func Test04ReadUnsafe(t *testing.T) {
	c := [200]byte{}
	q := &Queue{content: c[:]}
	v := q.ReadUnsafe(10)
	v[0] = 1

	fmt.Println(q.content[0]) // 1 q.content 值已经被修改
}

func Test04ReadSafe(t *testing.T) {
	c := [200]byte{}
	q := &Queue{content: c[:]}
	v := q.ReadSafe(10)
	v[0] = 1

	fmt.Println(q.content[0]) // 0 q.content safety
}
