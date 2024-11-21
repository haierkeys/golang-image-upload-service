FROM alpine:latest
ENV TZ=Asia/Shanghai
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories
RUN apk --update add libstdc++ curl ca-certificates bash curl gcompat tzdata && \
  cp /usr/share/zoneinfo/${TZ} /etc/localtime && \
  echo ${TZ} > /etc/timezone && \
  rm -rf  /tmp/* /var/cache/apk/*

EXPOSE 8000 8001 8002
RUN mkdir -p /api/
VOLUME /api/configs
VOLUME /api/storage
COPY ./build/linux/image-api /api/image-api
CMD ["sh", "-c","cd /api/ \
  && touch storage/logs/c.log \
  && mv storage/logs/c.log storage/logs/c.log_$(date '+%Y%m%d%H%M%S%'| cut -b 1-17) \
  && /api/image-api 2>&1 | tee storage/logs/c.log"]