build:
	go build -o playwright main.go

configure:
	glide install

fmt:
	go fmt

.PHONY: build configure fmt
