FROM windowsservercore
EXPOSE 80
COPY index.html .
COPY edit.html .
COPY create.html .
COPY goweb.exe C:\\app\\ .
ENTRYPOINT ["C:\\app\\goweb.exe"]
