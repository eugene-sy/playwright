GO_PATH=$(GOPATH)
GO_PATH?=/tmp/go
BINARY_NAME=playwright
REPO_PATH=$(CURDIR)
SRC_ROOT=$(REPO_PATH)/pkg
TEST_ROOT=$(REPO_PATH)/test

GO=go
GO_FMT=$(GO) fmt
GO_TEST=$(GO) test -v -p 1
GO_GET=$(GO) mod vendor
GO_VET=$(GO) vet
GO_INSTALL=$(GO) install

.PHONY: build
build:
	cd $(REPO_PATH) && go build -o $(BINARY_NAME) pkg/main.go

.PHONY: dependencies
dependencies:
	cd $(REPO_PATH) && $(GO_GET)

.PHONY: fmt
fmt:
	cd $(SRC_ROOT) && $(GO_FMT)
	cd $(SRC_ROOT)/commands && $(GO_FMT)
	cd $(SRC_ROOT)/utils && $(GO_FMT)
	cd $(SRC_ROOT)/logger && $(GO_FMT)

.PHONY: vet
vet:
	$(GO_VET) ./...

.PHONY: lint
lint:
	$(STATICCHECK) ./...

.PHONY: test
test:
	cd $(SRC_ROOT) && $(GO_TEST)
	cd $(SRC_ROOT)/commands && $(GO_TEST)
	cd $(SRC_ROOT)/utils && $(GO_TEST)
	cd $(SRC_ROOT)/logger && $(GO_TEST)

.PHONY: it
it:
	cd $(TEST_ROOT) && $(GO_TEST)

.PHONY: install-native
install-native: build
	cd $(REPO_PATH) && $(GO_INSTALL)

.PHONY: install
install: build
	cp $(BINARY_NAME) /usr/local/bin/$(BINARY_NAME)

.PHONY: clean
clean:
	rm -rf $(BINARY_NAME)

.PHONY: prepare
prepare:
	go install honnef.co/go/tools/cmd/staticcheck@latest
