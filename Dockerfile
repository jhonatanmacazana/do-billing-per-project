# first stage - builder
FROM golang:1.17 as builder

WORKDIR /app

# Copy dependency files
COPY go.mod go.sum ./

# Install dependencies
RUN go mod download

# Copy files
COPY . .

# Build project
# CGO_ENABLED=0
RUN GOOS=linux go build -a -ldflags '-extldflags "-static"' -o app ./cmd/do-billing-per-project/*



# second stage - main
FROM alpine:latest

# OS dependencies
RUN apk add --no-cache libc6-compat

WORKDIR /root/

COPY --from=builder /app .

CMD ["./app"]