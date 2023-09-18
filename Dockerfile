FROM golang:1.21.1-alpine
WORKDIR /usr/src/app
COPY go.mod go.mod
RUN go mod download && go mod verify
COPY . .
RUN go build -v -o /usr/local/bin/app ./cmd/main
CMD ["app"]