[English Document](README.md)

# Golang Image Upload Service

为 obsidian-auto-image-remote-uploader 提供图片上传/存储/同步云存储服务

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
mkdir -p /data/configs
mkdir -p /data/storage/logs
mkdir -p /data/storage/uploads

# 下载默认配置到配置文件目录内
wget https://raw.githubusercontent.com/haierspi/golang-image-upload-service/main/configs/config.yaml  -O /data/configs/config.yaml

# 创建&启动容器
docker run -tid --name image-api \
        -p 8000:8000 -p 8001:8001 -p 8002:8002 \
        -v /data/storage/logs/:/api/storage/logs/ \
        -v /data/storage/uploads/:/api/storage/uploads/ \
        -v /data/configs/:/api/configs/ \
        haierspi/golang-image-upload-service:latest

```

## 二进制下载安装

下载对应的服务器版本



## 图片上传API服务端

本插件需要配合**golang-image-upload-service** https://github.com/haierspi/golang-image-upload-service 才能正常使用

# 使用帮助

## 剪切板上传

支持黏贴剪切板的图片的时候直接上传，目前支持复制系统内图像直接上传。
支持通过设置 `frontmatter` 来控制单个文件的上传，默认值为 `true`，控制关闭请将该值设置为 `false`

支持 ".png", ".jpg", ".jpeg", ".bmp", ".gif", ".svg", ".tiff"

```yaml
---
image-auto-upload: true
---
```

## 批量上传一个文件中的所有图像文件

输入 `ctrl+P` 呼出面板，输入 `upload all images`，点击回车，就会自动开始上传。

路径解析优先级，会依次按照优先级查找：

1. 绝对路径，指基于库的绝对路径
2. 相对路径，以./或../开头
3. 尽可能简短的形式

## 批量下载网络图片到本地

输入 `ctrl+P` 呼出面板，输入 `download all images`，点击回车，就会自动开始下载。

## 支持右键菜单上传图片

目前已支持标准 md 以及 wiki 格式。支持相对路径以及绝对路径，需要进行正确设置，不然会引发奇怪的问题。

## 支持拖拽上传
支持图片的各种拖拽

## 感谢

https://github.com/renmu123/obsidian-image-auto-upload-plugin