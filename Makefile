test:
	go test -v ./tests/...
.PHONY: clean

make build:
	go build -o slim-go

make run:
	./slim-go
.PHONY: run
