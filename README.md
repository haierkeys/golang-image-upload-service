[中文文档](readme-zh.md) / [English Document](README.md)

# Obsidian Image API Gateway
===

Provides image upload/storage/sync services for the `obsidian-auto-image-remote-uploader` plugin.

Feature List:

- [x] Support for image uploads
- [x] Authorization token support for enhanced API security
- [x] HTTP access to images (basic functionality; using Nginx is recommended)
- [x] Storage-related features:
  - [x] Support for simultaneous storage on both local and cloud storage for easy migration
  - [x] Local storage support (prepared for NAS use,functionality supported and tested successfully)
  - [x] Support for Alibaba Cloud OSS (functionality supported but not yet tested)
  - [x] Support for Cloudflare R2 (functionality supported and tested successfully)
  - [x] Support for Amazon S3 (functionality supported and tested successfully)
  - [ ] Support for Google ECS (under development)
- [x] Docker installation for easy deployment on home NAS or remote servers
- [ ] Public API for users who are unable to set up their own API services

## Changelog

[Changelog](https://github.com/haierkeys/obsidian-image-api-gateway/releases)

## Pricing

This software is open source and free to use. However, if you'd like to show your support or help with continued development, feel free to contribute in any of the following ways:

- [![Paypal](https://img.shields.io/badge/paypal-haierkeys-yellow?style=social&logo=paypal)](https://paypal.me/haierkeys)

- [<img src="https://cdn.buymeacoffee.com/buttons/v2/default-yellow.png" alt="BuyMeACoffee" width="100">](https://www.buymeacoffee.com/haierkeys)

- <img src="https://raw.githubusercontent.com/haierkeys/obsidian-auto-image-remote-uploader/main/bmc_qr.png" style="width:120px;height:auto;">

- Afdian: https://afdian.net/a/haierkeys

# Getting Started

## Containerized Installation (Docker)

Assume your server’s image storage path is _/data/storage/uploads_.

Run the following commands:

```bash
# Pull the container image
docker pull haierkeys/obsidian-image-api-gateway:latest

# Create necessary directories for the project
mkdir -p /data/image-api/config
mkdir -p /data/image-api/storage/logs
mkdir -p /data/image-api/storage/uploads

# Download the default configuration file into the configuration directory
wget https://raw.githubusercontent.com/haierkeys/obsidian-image-api-gateway/main/configs/config.yaml  -O /data/config/config.yaml

# Create and start the container
docker run -tid --name image-api \
        -p 8000:8000 -p 8001:8001 \
        -v /data/image-api/storage/logs/:/api/storage/logs/ \
        -v /data/image-api/storage/uploads/:/api/storage/uploads/ \
        -v /data/image-api/config/:/api/config/ \
        haierkeys/obsidian-image-api-gateway:latest
```

## Binary Installation

Download the latest release from [GitHub Releases](https://github.com/haierkeys/obsidian-image-api-gateway/releases).

Extract it to the desired directory and execute the binary.

## Configuration

The default configuration file name is `_config.yaml_`, which should be located in the _root directory_ or the _config_ directory.

For detailed configuration instructions, refer to:

- [Configuration File with English Comments](config\config-en.yaml)


## TODO

## Others

**Obsidian Auto Image Remote Uploader**
[https://github.com/haierkeys/obsidian-auto-image-remote-uploader](https://github.com/haierkeys/obsidian-auto-image-remote-uploader)
