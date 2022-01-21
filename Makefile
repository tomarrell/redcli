.PHONY: build
build:
	go build ./cmd/redcli/...

.PHONY: run
run:
	go run ./cmd/redcli/...
