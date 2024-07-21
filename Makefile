entrypoint=./cmd/sylphy
tmp_bin=./tmp/bin/sylphy

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
	@SYLPHY_LOG_LEVEL=debug go run $(entrypoint)

.PHONY: test
test:
	@go test ./... -race

.PHONY: install
install:
	@echo "Installing Sylphy..."
	@go install $(entrypoint)
	@echo "Success"

.PHONY: clean
clean:
	rm $(tmp_bin)
	rm -rf ./config
