FROM golang:1.13 AS builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ./bin/scheduler ./cmd/scheduler/scheduler.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /app
RUN mkdir config bin logs
COPY --from=builder /app/bin/scheduler ./bin/scheduler
COPY  config/config.docker.yaml ./config/config.yaml
CMD ["./bin/scheduler"]
