################################################################################
# Run tests
#
FROM golang:latest AS test
COPY . /src/action/
WORKDIR /src/action/
RUN go test ./...  2>&1

################################################################################
# Build
#
FROM golang:latest AS build
COPY . /src/action/
WORKDIR /src/action/
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags="-w -s" -o /go/bin/action 2>&1

################################################################################
# Get CA-Certificates
#
FROM alpine:latest AS certs
RUN apk add --no-cache ca-certificates

################################################################################
# Build Image
#
FROM scratch
COPY --from=build /go/bin/action /action
COPY --from=certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
ENTRYPOINT ["/action"]

