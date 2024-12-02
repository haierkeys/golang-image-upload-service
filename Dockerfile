FROM alpine:latest
ARG TARGETOS
ARG TARGETARCH
ENV TZ=Asia/Shanghai
ENV P_NAME=api
ENV P_BIN=image-api
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories
RUN apk --update add libstdc++ curl ca-certificates bash curl gcompat tzdata && \
    cp /usr/share/zoneinfo/${TZ} /etc/localtime && \
    echo ${TZ} > /etc/timezone && \
    rm -rf  /tmp/* /var/cache/apk/*

EXPOSE 8000 8001 8002
RUN mkdir -p /api/
VOLUME /${P_NAME}/configs
VOLUME /${P_NAME}/storage
COPY ./build/${TARGETOS}_${TARGETARCH}/${P_BIN} /${P_NAME}/
CMD ["sh", "-c","cd /${P_NAME}/ \
    && mkdir -p storage/logs \
    && touch storage/logs/c.log \
    && mv storage/logs/c.log storage/logs/c.log_$(date '+%Y%m%d%H%M%S%'| cut -b 1-17) \
    && /${P_NAME}/${P_BIN} run 2>&1 | tee storage/logs/c.log"]