FROM mcr.microsoft.com/windows
EXPOSE 80
COPY templates .
COPY se.exe .
ENTRYPOINT ["C:\se.exe"]
