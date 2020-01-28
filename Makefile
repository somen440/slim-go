test:
	go test -v ./tests/...
.PHONY: clean

build:
	go build -o slim-go

run:
	./slim-go -addr :3333
.PHONY: run
