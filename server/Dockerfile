FROM alpine:3.12.0

COPY server.exe /app/server.exe
RUN chmod +x /app/server.exe \ 
 && ls -alF /app

WORKDIR /app
EXPOSE 3000

CMD ["./server.exe"]