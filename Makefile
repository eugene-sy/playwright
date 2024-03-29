BINARY_NAME=playwright
TEST_DIR=test
TEST_ROOT=./$(TEST_DIR)

GO=go
GO_BUILD=$(GO) build
GO_FMT=$(GO) fmt
GO_TEST=$(GO) test -v -p 1 -cover
GO_GET=$(GO) mod vendor
GO_VET=$(GO) vet
GO_INSTALL=$(GO) install
GO_LIST=$(GO) list
STATICCHECK=staticcheck

.PHONY: build
build:
	$(GO_BUILD) -o $(BINARY_NAME) pkg/main.go

.PHONY: dependencies
dependencies:
	$(GO_GET)

.PHONY: fmt
fmt:
	$(GO_FMT) ./...

.PHONY: vet
vet:
	$(GO_VET) ./...

.PHONY: lint
lint:
	$(STATICCHECK) ./...

.PHONY: test
test:
	$(GO_TEST) $$($(GO_LIST) ./... | grep -v $(TEST_DIR))

.PHONY: it
it:
	cd $(TEST_ROOT) && $(GO_TEST)

.PHONY: install-native
install-native: build
	$(GO_INSTALL)

.PHONY: install
install: build
	cp $(BINARY_NAME) /usr/local/bin/$(BINARY_NAME)

.PHONY: clean
clean:
	rm -rf $(BINARY_NAME)

.PHONY: prepare
prepare:
	go install honnef.co/go/tools/cmd/staticcheck@latest
