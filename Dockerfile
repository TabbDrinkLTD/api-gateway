FROM golang:1.9.0 as builder

WORKDIR /go/src/github.com/TabbDrinkLTD/api-gateway

COPY . .

RUN go get -u github.com/golang/dep/cmd/dep
RUN dep ensure

RUN CGO_ENABLED=0 GOOS=linux go build -o micro -a -installsuffix cgo .


FROM alpine:latest

RUN apk --no-cache add ca-certificates

RUN mkdir /micro
WORKDIR /micro
COPY --from=builder /go/src/github.com/TabbDrinkLTD/api-gateway/micro .

CMD ["./micro"]
