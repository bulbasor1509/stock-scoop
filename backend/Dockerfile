# Use Go base image for build and runtime
FROM golang:latest

# Set working directory
WORKDIR /app

# Set ROD Chromium path
ENV ROD_BROWSER_PATH=/usr/bin/chromium

# Copy Go mod files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire project
COPY . .

# Build the Go app from specific entry point
RUN go build -o main ./cmd/stock-scoop/main.go

# Install Chromium and required libraries
RUN apt-get update && apt-get install -y \
    chromium \
    fonts-liberation \
    libappindicator3-1 \
    libasound2 \
    libatk-bridge2.0-0 \
    libatk1.0-0 \
    libcups2 \
    libdbus-1-3 \
    libgdk-pixbuf2.0-0 \
    libnspr4 \
    libnss3 \
    libx11-xcb1 \
    libxcomposite1 \
    libxdamage1 \
    libxrandr2 \
    xdg-helpers \
    wget \
    ca-certificates \
    && apt-get clean && rm -rf /var/lib/apt/lists/*

# Expose application port
EXPOSE 8080

# Set the binary as entrypoint
ENTRYPOINT ["./main"]
