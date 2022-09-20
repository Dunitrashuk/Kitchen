FROM golang:1.17.1-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./
COPY /config ./config
COPY /structs ./structs
RUN go build -o /Kitchen

EXPOSE 8081
CMD ["/Kitchen"]