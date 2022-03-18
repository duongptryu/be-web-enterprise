FROM alpine
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
RUN update-ca-certificates
ADD https://github.com/golang/go/raw/master/lib/time/zoneinfo.zip /zoneinfo.zip
ENV ZONEINFO /zoneinfo.zip
ENV TZ="Asia/Ho_Chi_Minh"
RUN date

WORKDIR /app/
RUN mkdir -p /app/assets
ADD ./app /app/
ADD ./config_stg.yaml /app/config.yaml

ENTRYPOINT ["./app"]