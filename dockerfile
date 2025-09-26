FROM windowsservercore
EXPOSE 80
COPY templates .
COPY se.exe .
ENTRYPOINT ["C:\se.exe"]
