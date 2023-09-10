FROM alpine:latest
ENV TZ Asia/Shanghai
RUN apk --update add libstdc++ curl ca-certificates bash curl gcompat tzdata && \
      cp /usr/share/zoneinfo/${TZ} /etc/localtime && \
        echo ${TZ} > /etc/timezone && \
            rm -rf  /tmp/* /var/cache/apk/*

EXPOSE 8000 8001 8002
RUN mkdir -p /api/
VOLUME /api/configs
VOLUME /api/storage
COPY ./build/apiRun /api/apiRun
CMD ["sh", "-c","cd /api/ \
    && touch storage/logs/c.log \
    && mv storage/logs/c.log storage/logs/c.log_$(date '+%Y%m%d%H%M%S%'| cut -b 1-17) \
    && /api/apiRun 2>&1 | tee storage/logs/c.log"]