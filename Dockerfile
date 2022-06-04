FROM golang:1.17.7-alpine3.15
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go build -o simpleServer .
CMD ["/app/simpleServer"]