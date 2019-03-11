GOPATH?=$(HOME)/go
BIN?=./bin/main
LOCAL_BIN:=$(CURDIR)/bin

# Check global GOLANGCI-LINT
GOLANGCI_BIN:=$(LOCAL_BIN)/golangci-lint
GOLANGCI_TAG:=1.13.1

export GO111MODULE=on

# install golangci-lint binary
.PHONY: install-lint
install-lint:
ifeq ($(wildcard $(GOLANGCI_BIN)),)
	$(info #Downloading golangci-lint v$(GOLANGCI_TAG))
	go get -d github.com/golangci/golangci-lint@v$(GOLANGCI_TAG)
	go build -ldflags "-X 'main.version=$(GOLANGCI_TAG)' -X 'main.commit=test' -X 'main.date=test'" -o $(LOCAL_BIN)/golangci-lint github.com/golangci/golangci-lint/cmd/golangci-lint
	go mod tidy
GOLANGCI_BIN:=$(LOCAL_BIN)/golangci-lint
endif

# run full lint like in pipeline
.PHONY: lint
lint: install-lint
	$(GOLANGCI_BIN) run --config=.golangci.yaml ./...

.PHONY: .deps
.deps:
	$(info #Install dependencies...)
	go mod download

# install project dependencies
.PHONY: deps
deps: .deps


.PHONY: .build
.build:
	$(info #Building...)
	$(BUILD_ENVPARMS) go build -o $(BIN) ./main.go

# build app
.PHONY: build
build: .build

.PHONY: .run
.run:
	$(info #Running...)
	$(BUILD_ENVPARMS) go run ./main.go --local-config-enabled

# run app
.PHONY: run
run: .run


.PHONY: .test
.test:
	$(info #Running tests...)
	go test ./...

# run unit tests
.PHONY: test
test: .test