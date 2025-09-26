
FROM ubuntu
MAINTAINER js18.user@gmail.com
COPY se.exe .
EXPOSE 80
CMD ["se.exe", ]
