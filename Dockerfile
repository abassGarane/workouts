FROM golang:1.20rc2-bullseye
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . .
RUN go build -ldflags "-w -s" -o muscles cmd/web/*
CMD["/muscles"]