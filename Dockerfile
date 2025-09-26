FROM alpine
EXPOSE 80
COPY templates .
COPY se.exe .
# ENTRYPOINT ["se.exe"]
CMD ["./se.exe"]
