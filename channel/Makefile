all: clean
	go build -gcflags "-N -l" -ldflags=-compressdwarf=false -o 0-1-channel-close.out 0-1-channel-close.go
	go tool objdump -s "main.main" 0-1-channel-close.out > 0-1-channel-close.S
clean:
	rm -rf 0-1-channel-close.out 0-1-channel-close.S

default: test

test:
			@echo "[Test] running tests"
			@go test -v -count=1 -failfast .

bench:
			@echo "[Benchmark] running Benchmark"
			@go test -v -count=1 -bench=".*"
