FROM golang:alpine AS builder
WORKDIR /build
ADD go.mod .
COPY index.html .
COPY edit.html .
COPY create.html .
COPY goweb.go .

RUN go build -o hello goweb.go

FROM alpine

WORKDIR /build

COPY --from=builder /build/goweb /build/goweb

CMD [". /hello"]
