.PHONY: build run stop


build:
	go build -o build/app ./cmd/

run:
	nohup build/app > build/output.log 2>&1 & echo $$! > build/pid

stop:
	-kill `cat build/pid` && rm -f build/pid