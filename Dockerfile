FROM golang:1.14-alpine

ARG HOSTNAME=github.com
ARG NAMESPACE=farfetch-internal
ARG NAME=ffmsgraph
ARG BINARY=terraform-provider-${NAME}
ARG VERSION=1.0.0

WORKDIR /go/src/

RUN apk update \
    && apk add --no-cache git \
    && apk add --no-cache build-base

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/linux_amd64/${BINARY}

# copy /go/src/plugins/ into /root/.terraform.d/plugins/ at terraform base image for usage