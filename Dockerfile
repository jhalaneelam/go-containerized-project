FROM golang:1.18-alpine
RUN mkdir /app
ADD . /app
WORKDIR /app
COPY go.mod go.sum ./
RUN go build -o main .
ENTRYPOINT ["/app/main"]