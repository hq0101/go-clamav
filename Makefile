.PHONY: fmt
fmt: ## Run go fmt against code.
	go fmt ./...

.PHONY: run
run: ## Run a controller from your host.
	go run ./cmd/clamd-api/main.go

# 定义应用名称和默认版本号
APP1 = clamd-ctl
APP2 = clamd-api
DEFAULT_VERSION = v0.1

# 获取Git描述的版本号，如果未定义则使用默认值
# VERSION ?= $(shell git describe --tags --always --dirty 2>/dev/null || echo $(DEFAULT_VERSION))
VERSION ?= $(shell echo $(DEFAULT_VERSION))
# 定义目标平台和架构
PLATFORMS = darwin-amd64 darwin-arm64 linux-amd64 linux-arm64 windows-amd64 windows-arm64

# 输出目录
OUTPUT_DIR = bin

# 默认目标
.PHONY: build-all build-ctl build-api clean

# 构建所有目标
build-all: build-ctl build-api

# 构建cland-ctl
build-ctl: $(foreach platform,$(PLATFORMS),$(OUTPUT_DIR)/$(APP1)-$(VERSION).$(platform))

$(OUTPUT_DIR)/$(APP1)-$(VERSION).%: GOOS = $(word 1, $(subst -, ,$*))
$(OUTPUT_DIR)/$(APP1)-$(VERSION).%: GOARCH = $(word 2, $(subst -, ,$*))
$(OUTPUT_DIR)/$(APP1)-$(VERSION).%:
	@mkdir -p $(OUTPUT_DIR)
	GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o $@ cmd/$(APP1)/main.go
	@if [ "$(GOOS)" = "windows" ]; then mv $@ $@.exe; fi

# 构建cland-api
build-api: $(foreach platform,$(PLATFORMS),$(OUTPUT_DIR)/$(APP2)-$(VERSION).$(platform))

$(OUTPUT_DIR)/$(APP2)-$(VERSION).%: GOOS = $(word 1, $(subst -, ,$*))
$(OUTPUT_DIR)/$(APP2)-$(VERSION).%: GOARCH = $(word 2, $(subst -, ,$*))
$(OUTPUT_DIR)/$(APP2)-$(VERSION).%:
	@mkdir -p $(OUTPUT_DIR)
	GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o $@ cmd/$(APP2)/main.go
	@if [ "$(GOOS)" = "windows" ]; then mv $@ $@.exe; fi

# 清理目标文件
clean:
	rm -rf $(OUTPUT_DIR)

CONTAINER_TOOL ?= docker
IMG ?= clamd-api:latest

.PHONY: docker-build
docker-build: ## Build docker image with the manager.
	$(CONTAINER_TOOL) build -t ${IMG} .

.PHONY: docker-push
docker-push: ## Push docker image with the manager.
	$(CONTAINER_TOOL) push ${IMG}
