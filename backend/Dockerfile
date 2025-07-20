# Use Go image for both build and runtime to avoid GLIBC version mismatch
FROM golang:latest

WORKDIR /app

ENV ROD_BROWSER_PATH=/usr/bin/chromium

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main .

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

EXPOSE 8080

ENTRYPOINT ["./main"]
