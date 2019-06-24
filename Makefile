GO_PATH=$(GOPATH)
GO_PATH?=/tmp/go
BUILD_ROOT_PATH=$(GO_PATH)/src/github.com/Axblade
BINARY_NAME=playwright
REPO_PATH=$(CURDIR)

GO=go
GO_FMT=$(GO) fmt
GO_TEST=$(GO) test -v -p 1
GO_GET=$(GO) get
GO_INSTALL=$(GO) install

build:
	cd $(REPO_PATH) && go build -o $(BINARY_NAME) main.go

configure:
	cd $(REPO_PATH) && $(GO_GET)

fmt:
	cd $(REPO_PATH) && $(GO_FMT)
	cd $(REPO_PATH)/commands && $(GO_FMT)
	cd $(REPO_PATH)/utils && $(GO_FMT)
	cd $(REPO_PATH)/logger && $(GO_FMT)

test:
	cd $(REPO_PATH) && $(GO_TEST)
	cd $(REPO_PATH)/commands && $(GO_TEST)
	cd $(REPO_PATH)/utils && $(GO_TEST)
	cd $(REPO_PATH)/logger && $(GO_TEST)

install-native: build
	cd $(REPO_PATH) && $(GO_INSTALL)

install: build
	cp $(BINARY_NAME) /usr/local/bin/$(BINARY_NAME)

clean:
	rm -rf $(BINARY_NAME)

.PHONY: build configure fmt install install-native test
