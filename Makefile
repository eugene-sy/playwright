GO_PATH=$(GOPATH)
GO_PATH?=/tmp/go
BINARY_NAME=playwright
SRC_ROOT=$(BUILD_PATH)/pkg
TEST_ROOT=$(BUILD_PATH)/test
REPO_PATH=$(CURDIR)
VENDOR=$(REPO_PATH)/vendor

GO=go
GO_FMT=$(GO) fmt
GO_TEST=$(GO) test -v -p 1
GO_GET=$(GO) get
GO_INSTALL=$(GO) install

build:
	cd $(REPO_PATH) && go build -o $(BINARY_NAME) pkg/main.go

configure:
	cd $(REPO_PATH) && $(GO_GET)

fmt:
	cd $(SRC_ROOT) && $(GO_FMT)
	cd $(SRC_ROOT)/commands && $(GO_FMT)
	cd $(SRC_ROOT)/utils && $(GO_FMT)
	cd $(SRC_ROOT)/logger && $(GO_FMT)

test:
	cd $(SRC_ROOT) && $(GO_TEST)
	cd $(SRC_ROOT)/commands && $(GO_TEST)
	cd $(SRC_ROOT)/utils && $(GO_TEST)
	cd $(SRC_ROOT)/logger && $(GO_TEST)

it:
	cd $(TEST_ROOT) && $(GO_TEST)

install-native: build
	cd $(REPO_PATH) && $(GO_INSTALL)

install: build
	cp $(BINARY_NAME) /usr/local/bin/$(BINARY_NAME)

clean:
	rm -rf $(BINARY_NAME)

.PHONY: build configure fmt install install-native test it
