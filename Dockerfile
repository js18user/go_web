FROM alpine
# WORKDIR /app
EXPOSE 8080
COPY templates .
COPY se.exe .
CMD ["se.exe"]
