all: clean
	go build -gcflags "-N -l" -ldflags=-compressdwarf=false -o channel_close.out channel_close.go
	go tool objdump -s "main.main" channel_close.out > channel_close.S
clean:
	rm -rf channel_close.out channel_close.S
