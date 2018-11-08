default: test

test: 
			@echo "[Test] running tests"
			@go test -v -count=1 -failfast .

bench:
			@echo "[Benchmark] running Benchmark"
			@go test -v -count=1 -bench=".*"
