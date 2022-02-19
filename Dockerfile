#!/usr/bin/env bash
FROM alpine
FROM --platform=linux/amd64 BASE_IMAGE:VERSION
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
RUN update-ca-certificates

WORKDIR /app/
ADD ./app /app/
ADD ./config_stg.yaml /app/config.yaml
ENTRYPOINT ["./app"]