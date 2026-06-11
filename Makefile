.PHONY: build run test clean install dist lint vet docker-up docker-down docker-logs docker-restart

# Build variables
VERSION ?= $(shell git describe --tags --always --dirty 2>/dev/null || echo "dev")
LDFLAGS = -X main.version=$(VERSION)
BINARY = oc-go-cc
CMD = ./cmd/oc-go-cc

# ── Development ────────────────────────────────────────────────────

build:
	go build -ldflags "$(LDFLAGS)" -o bin/$(BINARY) $(CMD)

run:
	go run -ldflags "$(LDFLAGS)" $(CMD)

test:
	go test ./... -v -race

vet:
	go vet ./...

lint:
	@which golangci-lint > /dev/null || (echo "golangci-lint not found, please install it: https://golangci-lint.run/usage/install/" && exit 1)
	@echo "Running gofmt..."
	@test -z "$$(gofmt -d . | tee /dev/stderr)" || (echo "gofmt check failed" && exit 1)
	@echo "Running golangci-lint..."
	golangci-lint run --timeout 5m

clean:
	rm -rf bin/ dist/

install: build
	cp bin/$(BINARY) $(GOPATH)/bin/$(BINARY) 2>/dev/null || \
		cp bin/$(BINARY) $(HOME)/go/bin/$(BINARY) 2>/dev/null || \
		go install -ldflags "$(LDFLAGS)" $(CMD)

# ── Docker (compose) ────────────────────────────────────────────────

docker-up:
	@if [ ! -f .env ]; then echo "ERROR: .env not found. Create it with: cp .env.example .env"; exit 1; fi
	@mkdir -p .tmp/tiktoken-cache
	@if [ ! -f .tmp/tiktoken-cache/9b5ad71b2ce5302211f9c61530b329a4922fc6a4 ]; then \
		echo "Downloading tiktoken encoding..."; \
		wget -q -O .tmp/tiktoken-cache/9b5ad71b2ce5302211f9c61530b329a4922fc6a4 \
			https://openaipublic.blob.core.windows.net/encodings/cl100k_base.tiktoken; \
	fi
	DOCKER_BUILDKIT=1 docker compose up -d --build
	@echo ""
	@echo "Container started! Proxy listening on http://localhost:3456"
	@echo "Logs:      make docker-logs"
	@echo "Stop:      make docker-down"
	@echo "Restart:   make docker-restart"

docker-down:
	docker compose down

docker-logs:
	docker compose logs -f

docker-restart:
	docker compose restart

# ── Release / Cross-Compilation ────────────────────────────────────

PLATFORMS = \
	darwin-amd64 \
	darwin-arm64 \
	linux-amd64 \
	linux-arm64 \
	windows-amd64 \
	windows-arm64

RELEASE_LDFLAGS = $(LDFLAGS) -s -w

dist: clean
	@mkdir -p dist
	@echo "Building release binaries (version: $(VERSION))..."
	@for platform in $(PLATFORMS); do \
		IFS='-' read -r GOOS GOARCH <<< "$$platform"; \
		EXT=""; \
		[ "$$GOOS" = "windows" ] && EXT=".exe"; \
		echo "  → $$GOOS/$$GOARCH"; \
		CGO_ENABLED=0 GOOS=$$GOOS GOARCH=$$GOARCH \
			go build -ldflags "$(RELEASE_LDFLAGS)" \
				-o "dist/$(BINARY)_$${platform}$${EXT}" \
				$(CMD); \
	done
	@echo ""
	@echo "Generating checksums..."
	@cd dist && sha256sum $(BINARY)_* > checksums.txt
	@echo ""
	@echo "Built binaries:"
	@ls -lh dist/
