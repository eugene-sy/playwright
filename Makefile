GO_PATH=$(GOPATH)
GO_PATH?=/tmp/go
BUILD_ROOT_PATH=$(GO_PATH)/src/github.com/Axblade
BINARY_NAME=playwright
BUILD_PATH=$(BUILD_ROOT_PATH)/$(BINARY_NAME)
REPO_PATH=$(CURDIR)

GO=go
GO_FMT=$(GO) fmt
GO_TEST=$(GO) test -v -p 1

build:
	cd $(BUILD_PATH) && go build -o $(BINARY_NAME) main.go

configure:
	rm -r $(BUILD_ROOT_PATH) || true
	mkdir -p $(BUILD_ROOT_PATH)
	ln -s $(REPO_PATH) $(BUILD_PATH) || true
	cd $(BUILD_PATH) && dep ensure

fmt:
	cd $(BUILD_PATH) && $(GO_FMT)
	cd $(BUILD_PATH)/commands && $(GO_FMT)
	cd $(BUILD_PATH)/utils && $(GO_FMT)
	cd $(BUILD_PATH)/logger && $(GO_FMT)

test:
	cd $(BUILD_PATH) && $(GO_TEST)
	cd $(BUILD_PATH)/commands && $(GO_TEST)
	cd $(BUILD_PATH)/utils && $(GO_TEST)
	cd $(BUILD_PATH)/logger && $(GO_TEST)

install-native: build
	cd $(BUILD_PATH) && go install

install: build
	cp $(BINARY_NAME) /usr/local/bin/$(BINARY_NAME)

clean:
	rm -rf $(BINARY_NAME)
	rm -rf $(SRC)

.PHONY: build configure fmt install install-native test
