FROM golang:latest AS test

COPY . $GOPATH/src/action/
WORKDIR $GOPATH/src/action/

RUN go test ./...

FROM golang:latest AS build

COPY . $GOPATH/src/action/
WORKDIR $GOPATH/src/action/

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags="-w -s" -o /go/bin/action

FROM alpine:latest as certs
RUN apk --no-cache add tzdata zip ca-certificates

FROM scratch
COPY --from=build /go/bin/action /action
COPY --from=certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
ENTRYPOINT ["/action"]


