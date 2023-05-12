FROM golang:1.19

WORKDIR /app

COPY ./ ./
RUN go mod download

RUN mkdir -p /app/tmp

RUN CGO_ENABLED=0 GOOS=linux go build -o /promotions-app

CMD ["/promotions-app"]