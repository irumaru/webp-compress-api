##
# builder
##

FROM golang:1.25.5-alpine3.23 AS builder

RUN apk --no-cache add libwebp-dev build-base pkgconfig

WORKDIR /app
COPY . .

ENV CGO_ENABLED=1
RUN go build -o main main.go

##
# app
##

FROM alpine:3.23.0

RUN apk --no-cache add libwebp

WORKDIR /app
COPY --from=builder /app/main /app/

CMD ["/app/main"]
EXPOSE 1323
