build: configure
	go build -o playwright main.go

configure:
	glide install

fmt:
	go fmt

install: build
	cp playwright /usr/local/bin/playwright

.PHONY: build configure fmt
