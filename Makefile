.PHONY=all build run

all:
	make build && make run

build:
	go build  ./cmd/

run:
	nohup cmd > output.log 2>&1 &