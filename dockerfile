# syntax=docker/dockerfile:1

FROM golang:alpine as builder

# Set destination for COPY
WORKDIR /app

#COPY go.mod ./
#COPY go.sum ./
#RUN go mod download
COPY . /app

RUN go mod download
RUN go mod vendor

RUN ls

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o bin/util


FROM alpine:latest as runner

COPY --from=builder /app/bin .
COPY --from=builder /app/config /config

ENTRYPOINT ["./util"]

# Optional:
# To bind to a TCP port, runtime parameters must be supplied to the docker command.
# But we can document in the Dockerfile what ports
# the application is going to listen on by default.
# https://docs.docker.com/engine/reference/builder/#expose
EXPOSE 8080

# Run
#CMD ["/docker-gs-ping"]