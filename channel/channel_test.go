package channel

import "testing"

/*
* 测试表明当向关闭的channel中写入数据时会触发painc操作
* panic 类型: runtime.plainError
* panic 值: "send on closed channel"
 */
func TestCloseWriteChanException(t *testing.T) {
	if err := CloseWriteChanException(); err != nil {
		t.Error(err)
	}
}

/*
* 当往关闭的channel中读数据时, 当读到第二个返回值为false时 就表示无数据输入了;
* v, ok := <- chan
 */
func TestCloseReadChanException(t *testing.T) {
	if err := CloseReadChanException(); err != nil {
		t.Error(err)
	}
}
