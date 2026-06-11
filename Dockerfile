FROM golang:1.25-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -ldflags "-s -w -X main.version=docker" -o /app/oc-go-cc ./cmd/oc-go-cc

# Pre-download tiktoken encoding so the container doesn't need network at startup.
ENV TIKTOKEN_CACHE_DIR=/tmp/tiktoken-cache
RUN mkdir -p $TIKTOKEN_CACHE_DIR && \
    wget -q -O $TIKTOKEN_CACHE_DIR/9b5ad71b2ce5302211f9c61530b329a4922fc6a4 \
        https://openaipublic.blob.core.windows.net/encodings/cl100k_base.tiktoken

FROM alpine:3.21
RUN apk add --no-cache ca-certificates tzdata wget && \
    addgroup -S appgroup && adduser -S appuser -G appgroup && \
    mkdir -p /etc/oc-go-cc /home/appuser/.config/oc-go-cc && \
    chown -R appuser:appgroup /etc/oc-go-cc /home/appuser/.config/oc-go-cc

COPY --from=builder /app/oc-go-cc /usr/local/bin/oc-go-cc
COPY --from=builder /app/configs/config.example.json /etc/oc-go-cc/config.json
COPY --from=builder /tmp/tiktoken-cache /home/appuser/.cache/oc-go-cc/tiktoken
RUN chown -R appuser:appgroup /etc/oc-go-cc /home/appuser/.cache

USER appuser

ENV OC_GO_CC_CONFIG=/etc/oc-go-cc/config.json
ENV OC_GO_CC_HOST=0.0.0.0
ENV TIKTOKEN_CACHE_DIR=/home/appuser/.cache/oc-go-cc/tiktoken

EXPOSE 3456

HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
  CMD wget -qO- http://localhost:3456/health || exit 1

ENTRYPOINT ["oc-go-cc", "serve"]
