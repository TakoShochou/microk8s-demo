FROM scratch

ARG EXE_FILE

COPY ${EXE_FILE} /app/app.exe

WORKDIR /app
EXPOSE 3000

CMD ["./app.exe"]