FROM golang:1.10 as builder
WORKDIR $GOPATH/src/github.com/PeppyS/api.peppysisay.com/

# Install dep
RUN go get -u github.com/golang/dep/...

# Copy code from host and compile
COPY Gopkg.toml Gopkg.lock ./
RUN dep ensure --vendor-only
COPY . ./
RUN go build -o /bin/api cmd/api/api.go

# Copy binary to debian and run
FROM debian:jessie-slim

# Need ca-certificates to make https requests from container
RUN apt-get update
RUN apt-get install -y ca-certificates

COPY --from=builder /bin/api /bin/api

# Start API
ENTRYPOINT /bin/api
