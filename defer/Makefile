all: clean
	go build -gcflags "-N -l" -ldflags=-compressdwarf=false -o 0-4-compilation-analysis.out 0-4-compilation-analysis.go
	#go tool objdump -s "main.main" 0-4-compilation-analysis.out >  0-4-compilation-analysis.S
	go tool objdump -s "main.a" 0-4-compilation-analysis.out >  0-4-compilation-analysis.S


clean:
	rm -rf 0-4-compilation-analysis.S 0-4-compilation-analysis.out
