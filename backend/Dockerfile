# Build stage 
ARG GO_VERSION=1.24.5
FROM golang:${GO_VERSION} AS build

# Let Go auto-fetch newer toolchains if needed (matches your go.mod)
ENV GOTOOLCHAIN=auto

WORKDIR /app

# Cache deps
COPY go.mod go.sum ./
RUN go mod download

# Copy code and build
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o server ./cmd/api

# Final stage 
FROM scratch
WORKDIR /app
COPY --from=build /app/server .
COPY --from=build /app/migrations ./migrations
EXPOSE 8080
ENTRYPOINT ["./server"]
