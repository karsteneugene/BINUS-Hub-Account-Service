ARG GO_VERSION=1.18

FROM golang:${GO_VERSION}-alpine AS builder

RUN apk update && apk add alpine-sdk git && rm -rf /var/cache/apk/*

RUN mkdir -p /BINUS-Hub-Account-Service
WORKDIR /BINUS-Hub-Account-Service

COPY go.mod .
COPY go.sum .
RUN go mod tidy

COPY . .
RUN go build -o ./main

FROM alpine:latest

RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*

# RUN mkdir -p /BINUS-Hub-Account-Service
# WORKDIR /BINUS-Hub-Account-Service
COPY --from=builder /BINUS-Hub-Account-Service/main .
#COPY --from=builder /BINUS-Hub-Account-Service/docs /svc-project/docs

EXPOSE 8080

ENTRYPOINT ["./main","prod"]

