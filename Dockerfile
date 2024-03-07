
# For golang server
FROM golang:1.17
# Creating work directry inside docker container
WORKDIR /app
# Copying files from this computer to docker container
COPY login /app
COPY static /app/static
# Running server
RUN go build -o app
EXPOSE 3001
CMD ["./app"]
