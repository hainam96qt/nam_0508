FROM golang:1.19 AS go_builder

WORKDIR /app
COPY . /app

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags "-s -w" -o run

FROM alpine:3.16

WORKDIR /app

COPY --from=go_builder ./app/run ./nam-0508

RUN chmod +x /app/nam-0508

CMD ["/app/nam-0508"]


#FROM golang:latest as builder
#
#WORKDIR /app
#
#COPY . .
#
#RUN CGO_ENABLED=0 GOOS=linux go build -ldflags "-s -w" -o app
#
#FROM alpine
#
#RUN mkdir -p /app/config
#
#VOLUME /app/config
#
#COPY --from=builder /app/app /
#
#CMD ["/app"]