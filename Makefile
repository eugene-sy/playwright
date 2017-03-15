GO_PATH=/tmp/go
BUILD_ROOT_PATH=$(GO_PATH)/src/com.github/axblade
BUILD_PATH=$(BUILD_ROOT_PATH)/playwright
REPO_PATH=$(CURDIR)

build:
	cd $(BUILD_PATH) && go build -o playwright main.go

configure:
	rm -r $(BUILD_ROOT_PATH) || true
	mkdir -p $(BUILD_ROOT_PATH)
	ln -s $(REPO_PATH) $(BUILD_PATH) || true
	cd $(BUILD_PATH) && glide install

fmt:
	cd $(BUILD_PATH) && go fmt
	cd $(BUILD_PATH)/commands && go fmt
	cd $(BUILD_PATH)/utils && go fmt

install: build
	cp playwright /usr/local/bin/playwright

.PHONY: build configure fmt install
