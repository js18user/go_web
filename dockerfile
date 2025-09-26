FROM windowsservercore
EXPOSE 80
COPY se.exe .
ENTRYPOINT ["C:\se.exe"]
