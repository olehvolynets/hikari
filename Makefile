entrypoint=./cmd/sylphy
tmp_bin=./tmp/bin/sylphy

.PHONY: default
default: build run

.PHONY: build
build:
	@go build -o $(tmp_bin) $(entrypoint)

.PHONY: run
run:
	@$(tmp_bin) -config sample_config.yml

.PHONY: run_sample
run_sample:
	@$(tmp_bin) -config sample_config.yml test/sample.json

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
