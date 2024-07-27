entrypoint=./cmd/hikari
tmp_bin=./tmp/bin/hikari

.PHONY: all
all: run

.PHONY: build
build:
	@go build -o $(tmp_bin) $(entrypoint)

.PHONY: run
run:
	@go run $(entrypoint)

.PHONY: debug
debug:
	@HIKARI_LOG_LEVEL=debug go run $(entrypoint)

.PHONY: test
test:
	@go test ./... -race

.PHONY: install
install:
	@echo "Installing Hikari..."
	@go install $(entrypoint)
	@echo "Success"

.PHONY: clean
clean:
	rm $(tmp_bin)
	rm -rf ./config
