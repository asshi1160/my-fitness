FROM golang:1.20-alpine

ENV TZ Asia/Tokyo

ADD ./app /app
WORKDIR /app

RUN go install github.com/cosmtrek/air@latest
RUN go mod download

CMD ["air"]
