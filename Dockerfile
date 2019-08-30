FROM golang:1.12.5 as builder
ENV GO111MODULE=on
WORKDIR $GOPATH/src/github.com/PeppyS/api.peppysisay.com/

# Copy code from host and compile
COPY go.mod .
COPY go.sum .
RUN go mod download

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
