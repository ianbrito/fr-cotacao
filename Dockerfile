FROM golang:1.22.5 AS build

WORKDIR /app

COPY . /app

RUN go mod download && \
    go mod verify && \
    go build -o api ./cmd/api/main.go

FROM golang:1.22.5

ENV APP_ENV='production'
ENV HTTP_PORT=80

WORKDIR /opt/quote

COPY --from=build /app/api /opt/quote/api

EXPOSE $HTTP_PORT

CMD ["/opt/quote/api"]
