FROM alpine
EXPOSE 80
COPY index.html .
COPY edit.html .
COPY create.html .
COPY goweb.exe .
CMD ["goweb"]
