FROM golang:1.22-alpine

# ENV 

WORKDIR /app

COPY ./ /app

RUN go mod download

EXPOSE 8080

RUN go build -o -v cmd.api/main.go
RUN ./main
