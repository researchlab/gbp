# tire 

➜  trie git:(master) ✗ go test -v
=== RUN   TestFilter
--- PASS: TestFilter (0.00s)
	trie_test.go:11:  |中 false 1
	trie_test.go:11:  | |国 true 2
	trie_test.go:11:  | | |演 false 1
	trie_test.go:11:  | | | |员 true 0
	trie_test.go:11:  | | |人 true 0
	trie_test.go:11:  |好 false 1
	trie_test.go:11:  | |演 false 1
	trie_test.go:11:  | | |员 true 0
	trie_test.go:131: 测试通过, 期望得到:中国演员   实际得到:中国演员
	trie_test.go:11:  |中 false 1
	trie_test.go:11:  | |国 true 3
	trie_test.go:11:  | | |演 false 1
	trie_test.go:11:  | | | |员 true 0
	trie_test.go:11:  | | |人 true 0
	trie_test.go:11:  | | |商 false 1
	trie_test.go:11:  | | | |人 true 0
	trie_test.go:11:  |好 false 1
	trie_test.go:11:  | |演 false 1
	trie_test.go:11:  | | |员 true 0
	trie_test.go:142: 测试通过, 期望得到:中国商人   实际得到:中国商人
PASS
ok  	github.com/researchlab/experiments/detection/trie	0.002s

