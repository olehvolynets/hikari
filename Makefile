entrypoint=./cmd/sylphy
tmp_bin=./tmp/bin/sylphy
pkl_config=./pkl/Config.pkl

.PHONY: default
default: run

.PHONY: build
build:
	@go build -o $(tmp_bin) $(entrypoint)

.PHONY: run
run:
	@go run $(entrypoint)

.PHONY: test
test:
	@go test ./... -race

.PHONY: install
install:
	@echo "Installing Sylphy..."
	@go install $(entrypoint)
	@echo "Success"

.PHONY: pklgen
pklgen:
	@pkl-gen-go $(pkl_config)

.PHONY: clean
clean:
	rm $(tmp_bin)
