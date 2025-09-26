FROM mcr.microsoft.com/windows/base:ltsc2022
EXPOSE 8080
COPY templates .
COPY se.exe .
ENTRYPOINT ["C:\se.exe"]
