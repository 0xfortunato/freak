FROM golang:alpine as builder

LABEL mantainer="Mateus Fortunato <mateus__fortunato@outlook.com"

# Install Git
RUN apk update && apk add --no-cache git

# Set work directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download all dependencies. Which will be cached if the go.mod and go.sum are not changed
RUN go mod download

# Copy the source from the current dir to the working dir inside the container
COPY . .

# Build the app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Start new stage from scratch
FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the pre-built binary from the previous stage
COPY --from=builder /app/main .
COPY --from=builder /app/.env .

# Expose port 8080 to the outside
EXPOSE 8080

# Run the executable
CMD ["./main"]
