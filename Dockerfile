# FROM golang:1.12.0-alpine3.9
FROM golang:1.19
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go build -o main .
CMD ["/app/main"]