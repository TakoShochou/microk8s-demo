FROM scratch
ARG EXE_FILE

WORKDIR /app
EXPOSE 3000
COPY ${EXE_FILE} /app/app.exe

CMD ["./app.exe"]