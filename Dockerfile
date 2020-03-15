FROM golang:1.12.4-alpine3.9 AS builder

WORKDIR /usr/src/app

RUN apk update && apk add --no-cache --virtual .build-deps build-base git

COPY . .

RUN go build -o build/local/blog -v cmd/main.go

FROM alpine:3.9

RUN apk --no-cache add ca-certificates

WORKDIR /usr/local/bin

COPY --from=builder /usr/src/app/build/local/blog .
COPY --from=builder /usr/src/app/creditcards-e312a-346e92c9176a.json .
RUN export GOOGLE_APPLICATION_CREDENTIALS="./creditcards-e312a-346e92c9176a.json"
CMD ["./blog"]
