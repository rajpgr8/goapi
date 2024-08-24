FROM golang:1.23-bookworm as builder

WORKDIR /app

COPY go.* ./
RUN go mod download

COPY . ./

RUN echo $(ls -lrt ./cmd/api/main.go)

RUN go build -v -o server ./cmd/api/main.go


# Deploy the application binary into a lean image
FROM debian:bookworm-slim
COPY --from=builder /app/server /app/server

RUN echo $(ls -lrt /app/server)

EXPOSE 8081
ENTRYPOINT ["/app/server"]