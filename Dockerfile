
FROM golang:1.17
WORKDIR /app
COPY login /app
COPY static /app/static
RUN go build -o app
EXPOSE 3001
CMD ["./app"]
