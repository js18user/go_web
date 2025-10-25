FROM golang:alpine AS builder
WORKDIR /build
ADD go.mod .
ADD go.sum .
RUN go mod download

COPY goweb.go .

RUN go build -o goweb goweb.go

FROM alpine

WORKDIR /build

COPY index.html .
COPY edit.html .
COPY create.html .

COPY --from=builder /build/goweb /build/goweb
EXPOSE 80
CMD ["./goweb"]
