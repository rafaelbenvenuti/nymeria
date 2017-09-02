FROM golang:alpine as builder

RUN apk add --no-cache git curl build-base
RUN go get github.com/revel/cmd/revel

COPY . /go/src/github.com/rafaelbenvenuti/nymeria

RUN revel build github.com/rafaelbenvenuti/nymeria /tmp/nymeria

FROM alpine

COPY --from=builder /tmp/nymeria/ /nymeria

EXPOSE 9000

ENTRYPOINT /nymeria/run.sh
