FROM mcr.microsoft.com/windows/servercore:ltsc2022
EXPOSE 80
COPY index.html .
COPY edit.html .
COPY create.html .
COPY goweb.exe C:\\app\\ .
ENTRYPOINT ["C:\\app\\goweb.exe"]
