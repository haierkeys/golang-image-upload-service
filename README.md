[中文文档](readme-zh.md)
# Golang Image Upload Service

Image upload/storage/synchronization service for the obsidian-auto-image-remote-uploader plugin.

- Support image upload
- support authorization tokens, increase API security
- image http access (basic function, suggest to use nginx instead)
- Synchronization cloud storage (AliCloud OSS, Amazon S3, Google ECS).
- Docker command installation, easy to use in the home NAS and websites
- No need to build an environment to download and run directly

## Price

This plugin is provided free of charge to everyone, but if you would like to show your appreciation or help support the continued development, please feel free to provide me with a little help in any of the following ways:

- [![Paypal](https://img.shields.io/badge/paypal-HaierSpi-yellow?style=social&logo=paypal)](https://paypal.me/haierspi)

- [<img src="https://cdn.buymeacoffee.com/buttons/v2/default-yellow.png" alt="BuyMeACoffee" width="100">](https://www.buymeacoffee.com/haierspi)
<img src="https://raw.githubusercontent.com/haierspi/obsidian-auto-image-remote-uploader/main/bmc_qr.png" style="width:120px;height:auto;">

- afdian: https://afdian.net/a/haierspi
# Getting Started

## Containerized installation (docker way)

Assuming your server image storage path is */data/storage/uploads*, run the following commands in order to install your server.

Execute the following commands in order

```bash
# Download the container image
docker pull haierspi/golang-image-upload-service:latest

# Create the necessary directories for the project to run
mkdir -p /data/configs
mkdir -p /data/storage/logs
mkdir -p /data/storage/uploads

# Download the default configuration to the config file directory
wget https://raw.githubusercontent.com/haierspi/golang-image-upload-service/main/configs/config.yaml -O /data/configs/config.yaml

# Create & start the container
docker run -tid --name image-api \
        -p 8000:8000 -p 8001:8001 \
        -v /data/storage/logs/:/api/storage/logs/ \
        -v /data/storage/uploads/:/api/storage/uploads/ \
        -v /data/configs/:/api/configs/ \
        haierspi/golang-image-upload-service:latest

```

## Binary download and installation

https://github.com/haierspi/golang-image-upload-service/releases Download the latest version

Unzip it to the appropriate directory and run
## Configuration

Configuration file path *. /configs/config.yaml*

The default content is as follows

```yaml
Server:
  RunMode: debug
  # Service ports in the form IP:PORT (destined to listen on IP) or :PORT (listen on all)
  HttpPort:  :8000
  ReadTimeout: 60
  WriteTimeout: 60
  # Performance Listening Interface
  PrivateHttpListen:  :8001
Security:
  # Image Upload API Authorization TOKEN
  AuthToken: 6666
App:
  DefaultPageSize: 10
  MaxPageSize: 100
  DefaultContextTimeout: 60
  LogSavePath: storage/logs
  LogFileName: app
  LogFileExt: .log
  # Image file storage path
  UploadSavePath: storage/uploads
  # Access address for uploading attachments, including UploadSavePath, which describes the URL prefix that the interface returns to the uploader.
  UploadServerUrl: http://127.0.0.1:8000/storage/uploads
  # Upload size limit; unit: MB
  UploadImageMaxSize: 5
  # Upload Image File types limit
  UploadImageAllowExts:
    - .jpg
    - .jpeg
    - .png
    - .bmp
    - .gif
    - .svg
    - .tiff
# AliCloud OSS
OSS:
  # Whether to enable OSS cloud storage
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

## Other

Obsidian Auto Image Remote Uploader

https://github.com/haierspi/obsidian-auto-image-remote-uploader