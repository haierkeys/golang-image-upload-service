[English Document](README.md)

# Golang Image Upload Service

为 obsidian-auto-image-remote-uploader 插件提供图片上传/存储/同步云存储服务

- 支持图片上传
- 支持授权令牌,增加API安全
- 图片http访问(基础功能,建议使用nginx替代)
- 同步云存储(阿里云OSS,亚马逊S3,Google ECS)
- Docker命令安装,方便大家使用在家庭NAS和网站中
- 不需要搭建环境下载下来就可以直接运行

## 价格

本软件是开源软件,免费提供给大家的使用，但如果您想表示感谢或帮助支持继续开发，请随时通过以下任一方式为我提供一点帮助：

- [![Paypal](https://img.shields.io/badge/paypal-HaierSpi-yellow?style=social&logo=paypal)](https://paypal.me/haierspi)

- [<img src="https://cdn.buymeacoffee.com/buttons/v2/default-yellow.png" alt="BuyMeACoffee" width="100">](https://www.buymeacoffee.com/haierspi)
<img src="https://raw.githubusercontent.com/haierspi/obsidian-auto-image-remote-uploader/main/bmc_qr.png" style="width:120px;height:auto;">

- 爱发电 https://afdian.net/a/haierspi

# 开始使用

## 容器化安装 (docker方式)

假设您的服务器图片保存路径为 */data/storage/uploads*

依次执行以下命令

```bash
# 下载容器镜像
docker pull haierspi/golang-image-upload-service:latest

# 创建项目运行必要目录
mkdir -p /data/image-api/config
mkdir -p /data/image-api/storage/logs
mkdir -p /data/image-api/storage/uploads

# 下载默认配置到配置文件目录内
wget https://raw.githubusercontent.com/haierspi/golang-image-upload-service/main/configs/config.yaml  -O /data/config/config.yaml

# 创建&启动容器
docker run -tid --name image-api \
        -p 8000:8000 -p 8001:8001 \
        -v /data/image-api/storage/logs/:/api/storage/logs/ \
        -v /data/image-api/storage/uploads/:/api/storage/uploads/ \
        -v /data/image-api/config/:/api/config/ \
        haierspi/golang-image-upload-service:latest

```

## 二进制下载安装

https://github.com/haierspi/golang-image-upload-service/releases 下载最新版本

解压到相应的目录执行

## 配置

配置文件路径 *./configs/config.yaml*

默认内容如下

```yaml
Server:
  RunMode: debug
  # 服务端口 格式为 IP:PORT (注定监听IP) 或者 :PORT (监听所有)
  HttpPort:  :8000
  ReadTimeout: 60
  WriteTimeout: 60
  # 性能监听接口
  PrivateHttpListen:  :8001
Security:
  # 图片上传API授权TOKEN
  AuthToken: 6666
App:
  DefaultPageSize: 10
  MaxPageSize: 100
  DefaultContextTimeout: 60
  LogSavePath: storage/logs
  LogFileName: app
  LogFileExt: .log
  # 图片文件存储地址
  UploadSavePath: storage/uploads
  TempPath: storage/temp
  # 上传附件访问地址,需要包含 UploadSavePath, 这里用来描述接口返回给上传端使用的URL前缀
  UploadServerUrl: http://127.0.0.1:8000/storage/uploads
  # 上传大小限制 单位MB
  UploadImageMaxSize: 5
  # 上传图片类型限制
  UploadImageAllowExts:
    - .jpg
    - .jpeg
    - .png
    - .bmp
    - .gif
    - .svg
    - .tiff
# 阿里云OSS
OSS:
  # 是否开启OSS云存储
  Enable: false
  BucketName:
  Endpoint:
  AccessKeyID:
  AccessKeySecret:

Email:
  Host: smtp.gmail.com
  Port: 465
  UserName: xxx
  Password: xxx
  IsSSL: true
  From: xxx
  To:
    - xxx
```
## TODO

## 其他

Obsidian Auto Image Remote Uploader

https://github.com/haierspi/obsidian-auto-image-remote-uploader