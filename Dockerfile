FROM golang:1.23.6-alpine
WORKDIR /app
COPY index.html .
COPY edit.html .
COPY create.html .
COPY goweb.go .
RUN go build goweb.go
EXPOSE 80
CMD ["./goweb"]
