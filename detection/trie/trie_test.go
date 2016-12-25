package trie

import (
	"testing"
)

func printTrie(node *Node, t *testing.T, line string) {
	if len(node.Node) > 0 {
		for char, n := range node.Node {
			//t.Logf("%s%s", line, string(char))
			t.Logf("%s%s %t %d", line, string(char), n.End, len(n.Node))
			printTrie(n, t, line+" |")
		}
	}
}

//func TestAdd(t *testing.T) {
//	trie := NewTrie()
//	trie.Add("中华人民共和国")
//	trie.Add("中国")
//	trie.Add("中国人民")
//	trie.Add("中国共产党")
//	trie.Add("中国人民解放军")
//	trie.Add("中国人民武警")
//	trie.Add("中央书记")
//	trie.Add("华人")
//	trie.Add("我men是")
//
//	node := trie.Root
//	printTrie(node, t, " |")
//
//	words := trie.ReadAll()
//	for _, w := range words {
//		t.Logf("%s\n", w)
//	}
//}
//
//func TestDel(t *testing.T) {
//	trie := NewTrie()
//	trie.Add("AV")
//	trie.Add("AV演员")
//	trie.Add("AV演员色情")
//	trie.Add("日本AV女优")
//
//	text := "AV AV演员 AV演员色情"
//	expect := ""
//	got := ""
//
//	printTrie(trie.Root, t, " |")
//	t.Log("-----------------------")
//
//	//删除开头的
//	expect = "AV **** ******"
//	trie.Del("AV")
//	_, _, got = trie.Query(text)
//
//	if got != expect {
//		t.Errorf("希望得到: %s\n实际得到: %s\n", expect, got)
//	}
//	trie.Add("AV")
//
//	// 删除中间的
//	trie.Del("AV演员")
//	expect = "** **演员 ******"
//	_, _, got = trie.Query(text)
//	if got != expect {
//		t.Errorf("希望得到: %s\n实际得到: %s\n", expect, got)
//	}
//	trie.Add("AV演员")
//
//	// 删除后面的
//	trie.Del("AV演员色情")
//	expect = "** **** ****色情"
//	_, _, got = trie.Query(text)
//	if got != expect {
//		t.Errorf("希望得到: %s\n实际得到: %s\n", expect, got)
//	}
//	trie.Add("AV演员色情")
//
//	//删除不存在的敏感词
//	trie.Del("VA演")
//	expect = "** **** ******"
//	_, _, got = trie.Query(text)
//	if got != expect {
//		t.Errorf("希望得到: %s\n实际得到: %s\n", expect, got)
//	}
//
//	trie.Del("AV演员色情表演")
//	expect = "** **** ******"
//	_, _, got = trie.Query(text)
//	if got != expect {
//		t.Errorf("希望得到: %s\n实际得到: %s\n", expect, got)
//	}
//}
//
//func TestDel2(t *testing.T) {
//	trie := NewTrie()
//	trie.Add("中")
//
//	trie.Del("中")
//	words := trie.ReadAll()
//
//	if len(words) > 0 {
//		t.Error("只有一个字的删除失败")
//	}
//}

func TestFilter(t *testing.T) {
	trie := NewTrie()
	trie.Add("中国")
	trie.Add("中国演员")
	trie.Add("中国人")
	trie.Add("好演员")

	node := trie.Root
	printTrie(node, t, " |")
	text_list := []string{
		"在中国出生就是中国人，如果你是个演员，还是在中国出生，那就是中国演员，希望你努力成为一个好演员",
		"在中国出生就是中国人，如果你是个演员，还是在中国出生，那就是中国演员，希望你努力成为一个好演员,不过成为一个中国商人也蛮好的",
	}
	expect_list := []string{
		"中国演员",
		"中国商人",
	}
	//text := "在中国出生就是中国人，如果你是个演员，还是在中国出生，那就是中国演员，希望你努力成为一个好演员"
	//expect := "中国演员"
	max_len := trie.Filter(text_list[0])
	if max_len != expect_list[0] {
		t.Errorf("没有通过测试, 期望得到:%s   实际得到:%s\n", expect_list[0], max_len)
	} else {
		t.Logf("测试通过, 期望得到:%s   实际得到:%s\n", expect_list[0], max_len)
	}

	trie.Add("中国商人")
	node = trie.Root
	printTrie(node, t, " |")

	last_one := trie.Filter(text_list[1])
	if last_one != expect_list[1] {
		t.Errorf("没有通过测试, 期望得到:%s   实际得到:%s\n", expect_list[1], last_one)
	} else {
		t.Logf("测试通过, 期望得到:%s   实际得到:%s\n", expect_list[1], last_one)
	}
}

func BenchmarkFilter_Max(b *testing.B) {
	trie := NewTrie()
	trie.Add("中国")
	trie.Add("中国演员")
	trie.Add("中国人")
	trie.Add("好演员")
	trie.Add("中国商人")
	text_list := []string{
		"在中国出生就是中国人，如果你是个演员，还是在中国出生，那就是中国演员，希望你努力成为一个好演员",
		"在中国出生就是中国人，如果你是个演员，还是在中国出生，那就是中国演员，希望你努力成为一个好演员,不过成为一个中国商人也蛮好的",
	}
	for i := 0; i < b.N; i++ {
		trie.Filter(text_list[0])
	}
}

func BenchmarkFilter_Last(b *testing.B) {
	trie := NewTrie()
	trie.Add("中国")
	trie.Add("中国演员")
	trie.Add("中国人")
	trie.Add("好演员")
	trie.Add("中国商人")
	text_list := []string{
		"在中国出生就是中国人，如果你是个演员，还是在中国出生，那就是中国演员，希望你努力成为一个好演员",
		"在中国出生就是中国人，如果你是个演员，还是在中国出生，那就是中国演员，希望你努力成为一个好演员,不过成为一个中国商人也蛮好的",
	}
	for i := 0; i < b.N; i++ {
		trie.Filter(text_list[1])
	}
}

//func TestQuery(t *testing.T) {
//	trie := NewTrie()
//	trie.Add("AV")
//	trie.Add("AV演员")
//	trie.Add("AV演员色情")
//	trie.Add("日本AV女优")
//
//	node := trie.Root
//	printTrie(node, t, " |")
//
//	text := "日本AV演员兼电视、电影演员。苍井空AV女优是xx出道, 日本AV女优们最精彩的表演是AV演员色情表演"
//	expect := "日本****兼电视、电影演员。苍井空**女优是xx出道, ******们最精彩的表演是******表演"
//
//	ok, words, newText := trie.Query(text)
//
//	t.Log("words:", words)
//	t.Log("text:", newText)
//
//	if !ok {
//		t.Error("替换失败 1")
//	}
//
//	if len(words) == 0 {
//		t.Error("替换失败 2")
//	}
//
//	if newText != expect {
//		t.Errorf("希望得到: %s\n实际得到: %s\n", expect, newText)
//	}
//
//	// 和谐的文本
//	text = "完全和谐的文本完全和谐的文本"
//	ok, words, newText = trie.Query(text)
//
//	if ok {
//		t.Error("替换失败")
//	}
//
//	if len(words) != 0 {
//		t.Error("替换失败 2")
//	}
//
//	if newText != text {
//		t.Error("替换失败 3")
//	}
//}
//
//func TestQuery2(t *testing.T) {
//	trie := NewTrie()
//	trie.Add("口交")
//	trie.Add("口交女")
//
//	node := trie.Root
//	printTrie(node, t, " |")
//
//	text := "XX路口交"
//
//	ok, words, newText := trie.Query(text)
//
//	t.Log("words:", words)
//	t.Log("text:", newText)
//
//	if !ok {
//		t.Error("替换失败")
//	}
//}
//
//func TestReplaceNilTrie(t *testing.T) {
//	trie := NewTrie()
//	// 和谐的文本
//	text := "完全和谐的文本完全和谐的文本"
//	ok, words, newText := trie.Query(text)
//
//	if ok {
//		t.Error("替换失败")
//	}
//
//	if len(words) != 0 {
//		t.Error("替换失败 2")
//	}
//
//	if newText != text {
//		t.Error("替换失败 3")
//	}
//}
