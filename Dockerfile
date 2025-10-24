FROM golang:1.23.6-alpine
WORKDIR /app
COPY index.html .
COPY edit.html .
COPY create.html .
COPY goweb.go .
RUN go build -c goweb.go
EXPOSE 80
CMD ["./goweb"]

FROM golang:alpine AS builder

WORKDIR /build

ADD go.mod .

COPY . .

RUN go build -o hello hello.go

FROM alpine

WORKDIR /build

COPY --from=builder /build/hello /build/hello

CMD [". /hello"]
