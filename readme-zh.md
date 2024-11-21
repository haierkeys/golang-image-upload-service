[English Document](README.md)

# Obsidian Image Api Gateway
===

为 obsidian-auto-image-remote-uploader 插件提供图片上传/存储/同步云存储服务

功能清单:

- [x] 支持图片上传
- [x] 支持授权令牌,增加 API 安全
- [x] 图片http访问 (基础功能,建议使用 nginx 替代)
- [x] 存储相关:
  - [x] 支持同时保存到本地或云存储,方便后续迁移
  - [x] 支持本地保存 ( 为 NAS 准备 )
  - [x] 支持 阿里云 OSS 云存储(功能支持,尚未测试)
  - [x] 支持 Cloudflare R2 云存储(功能支持,测试OK)
  - [ ] 支持 Amazon S3 (待开发)
  - [ ] 支持 Google ECS (待开发)
- [x] Docker命令安装,方便大家使用在家庭NAS和远端服务器中
- [ ] 增加公共API,针对不方便架设 API 服务的用户

## 变更日志
- v0.5
  增加对 AWS S3, Cloudflare R2 云存储的支持,
  增加 存储方式同时执行
  为了方便识别,将项目名改为 obsidian-image-api-gateway


## 价格

本软件是开源软件,免费提供给大家的使用，但如果您想表示感谢或帮助支持继续开发，请随时通过以下任一方式为我提供一点帮助：

- [![Paypal](https://img.shields.io/badge/paypal-haierkeys-yellow?style=social&logo=paypal)](https://paypal.me/haierkeys)

- [<img src="https://cdn.buymeacoffee.com/buttons/v2/default-yellow.png" alt="BuyMeACoffee" width="100">](https://www.buymeacoffee.com/haierkeys)

- <img src="https://raw.githubusercontent.com/haierkeys/obsidian-auto-image-remote-uploader/main/bmc_qr.png" style="width:120px;height:auto;">

- 爱发电 https://afdian.net/a/haierkeys

# 开始使用

## 容器化安装 (docker方式)

假设您的服务器图片保存路径为 _/data/storage/uploads_

依次执行以下命令

```bash

# 下载容器镜像
docker pull haierkeys/obsidian-image-api-gateway:latest

# 创建项目运行必要目录
mkdir -p /data/image-api/config
mkdir -p /data/image-api/storage/logs
mkdir -p /data/image-api/storage/uploads

# 下载默认配置到配置文件目录内
wget https://raw.githubusercontent.com/haierkeys/obsidian-image-api-gateway/main/configs/config.yaml  -O /data/config/config.yaml

# 创建&启动容器
docker run -tid --name image-api \
        -p 8000:8000 -p 8001:8001 \
        -v /data/image-api/storage/logs/:/api/storage/logs/ \
        -v /data/image-api/storage/uploads/:/api/storage/uploads/ \
        -v /data/image-api/config/:/api/config/ \
        haierkeys/obsidian-image-api-gateway:latest

```

## 二进制下载安装

https://github.com/haierkeys/obsidian-image-api-gateway/releases 下载最新版本

解压到相应的目录执行

## 配置

配置文件路径 _./configs/config.yaml_

默认内容如下

```yaml
server:
  run-mode:
  # 服务端口-格式为-ip:port (注定监听ip) 或者 :port (监听所有)
  http-port:  :8000
  read-timeout: 60
  write-timeout: 60
  # 性能监听接口
  private-http-listen:  :8001
security:
  # 图片上传api授权token
  auth-token: 6666

app:
  default-page-size: 10
  max-page-size: 100
  default-context-timeout: 60
  log-save-path: storage/logs
  log-file: app.log

  temp-path: storage/temp
  # 接口返回上传前缀地址
  upload-url-pre: https://image.diybeta.com
  # 上传大小限制-单位mb
  upload-max-size: 5
  # 上传图片类型限制
  upload-allow-exts:
    - .jpg
    - .jpeg
    - .png
    - .bmp
    - .gif
    - .svg
    - .tiff
    - .heif
    - .avif
    - .webp

# 本地存储
local-fs:
  enable: true
  # 是否启用内置的文件URL访问服务
  httpfs-enable: true
  save-path: storage/uploads


# 阿里云oss
oss:
  # 是否开启oss云存储
  enable: false
  custom-path: blog
  bucket-name:
  endpoint:
  access-key-id:
  access-key-secret:

cloudflu-r2:
  # 是否开启R2云存储
  enable: true
  custom-path: blog
  bucket-name: image
  account-id:
  access-key-id:
  access-key-secret:

email:
  error-report-enable: false
  host: smtp.gmail.com
  port: 465
  user-name: xxx
  password: xxx
  is-ssl: true
  from: xxx
  to:
    - xxx
```

## TODO

## 其他

Obsidian Auto Image Remote Uploader

https://github.com/haierkeys/obsidian-auto-image-remote-uploader
