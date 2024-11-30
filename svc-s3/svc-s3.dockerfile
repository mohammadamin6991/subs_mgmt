# base go image
FROM m.reg.amin.run/golang:1.23-alpine as builder

RUN mkdir /app
COPY ./ /app
WORKDIR /app

RUN CGO_ENABLED=0 go build -o ./bin/svcS3 ./cmd/api
RUN chmod +x /app/bin/svcS3

# build tiny docker image
FROM m.reg.amin.run/alpine:3

RUN mkdir /app
COPY --from=builder /app/bin/svcS3 /app

CMD ["/app/svcS3"]
