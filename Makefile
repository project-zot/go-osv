TOOLS_DIR := hack/tools
TOOLS_BIN_DIR := $(TOOLS_DIR)/bin
BIN_DIR := bin

GOLANGCI_LINT := $(TOOLS_BIN_DIR)/golangci-lint

.PHONY: all
all: build check binary

.PHONY: binary
binary:
	mkdir -p $(BIN_DIR)
	go build -v -gcflags all='-N -l' -o $(BIN_DIR)/osv ./cmd/osv/...

.PHONY: build
build:
	go build -v ./...

$(GOLANGCI_LINT):
	mkdir -p $(TOOLS_BIN_DIR)
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(TOOLS_BIN_DIR) v1.50.1

.PHONY: check
check: $(GOLANGCI_LINT)
	$(GOLANGCI_LINT) run -c golangcilint.yaml ./...

.PHONY: clean
clean:
	rm -rf $(BIN_DIR)
	rm -rf $(TOOLS_DIR)
