FROM alpine
WORKDIR /app
EXPOSE 80
COPY templates .
COPY se.exe .
CMD ["./se.exe"]
