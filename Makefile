build:
	go build -o playwright main.go

configure:
	glide install

.PHONY: build configure
