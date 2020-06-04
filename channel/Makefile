all: clean
	go build -gcflags "-N -l" -ldflags=-compressdwarf=false -o channel_close.out channel_close.go
	go tool objdump -s "main.main" channel_close.out > channel_close.S
clean:
	rm -rf channel_close.out channel_close.S

default: test

test:
			@echo "[Test] running tests"
			@go test -v -count=1 -failfast .

bench:
			@echo "[Benchmark] running Benchmark"
			@go test -v -count=1 -bench=".*"
