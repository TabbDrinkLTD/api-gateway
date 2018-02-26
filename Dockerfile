FROM alpine:latest

RUN apk --no-cache add ca-certificates

RUN mkdir /micro
WORKDIR /micro
COPY micro .

CMD ["./micro"]
