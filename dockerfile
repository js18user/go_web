# syntax=docker/dockerfile:2

FROM alpine

LABEL maintainer="Jurij <js18.user@gmail.com>"

RUN adduser --disabled-password --gecos '' appuser

USER appgo

COPY se.exe .

EXPOSE 80

CMD ["se.exe" ]
