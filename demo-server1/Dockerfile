FROM alpine:3.12.0

ARG EXE_FILE

COPY ${EXE_FILE} /app/app.exe
RUN chmod +x /app/app.exe \ 
 && ls -alF /app

WORKDIR /app
EXPOSE 3000

CMD ["./app.exe"]