#x64
FROM golang:alpine

ENV GO112MODULE=on

RUN mkdir /app
ADD . /app/
WORKDIR /app

COPY . .

#x64
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 #go build

RUN go build -o main .
RUN adduser -S -D -H -h /app appuser
USER appuser
CMD ["./main"]