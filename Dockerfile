# Stage 1: Build the Go application
FROM golang:1.19-alpine AS builder

WORKDIR /go/src/app

# Copy the Go modules files
COPY go.mod ./

# Download the dependencies
RUN go mod tidy

COPY main.go .

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /go/bin/app .

# Stage 2: Create a minimal container for running the application
FROM gcr.io/distroless/static-debian11

COPY --from=builder /go/bin/app /
CMD ["/app"]
