ARG GO_IMAGE=golang:1.23.12
FROM ${GO_IMAGE}

ENV GOTOOLCHAIN=local \
    CGO_ENABLED=0 \
    GOPROXY=https://goproxy.cn,direct

WORKDIR /workspace
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build ./...

CMD ["go", "test", "./...", "-count=1"]
