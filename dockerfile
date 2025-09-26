
FROM ubuntu
MAINTAINER js18.user@gmail.com
COPY se.exe /bin/
EXPOSE 80
CMD ["/bin/name.exe", ]
