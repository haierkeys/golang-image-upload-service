[中文文档](readme-zh.md)
[English Document](README.md)

# Obsidian Image API Gateway

### A Gateway Service for Image Upload/Storage/Sync with Cloud Storage for the `obsidian-auto-image-remote-uploader` Plugin

---

### Features:

- [x] **Image Upload Support**
- [x] **Authorization Tokens** for enhanced API security
- [x] **HTTP Access to Images** (basic feature; consider using Nginx as an alternative)
- [x] **Storage Options:**
  - [x] Save images locally and/or on cloud storage for easy migration
  - [x] Local storage support (ideal for NAS setups)
  - [x] Aliyun OSS Cloud Storage (functional, untested)
  - [x] Cloudflare R2 Cloud Storage (functional, tested)
  - [ ] Amazon S3 Support (under development)
  - [ ] Google ECS Support (under development)
- [x] **Docker Installation** for convenient deployment on home NAS or remote servers
- [ ] **Public API** for users unable to host their own API service

---

## Changelog

### v0.5

- Added support for AWS S3 and Cloudflare R2 storage.
- Introduced simultaneous execution of multiple storage methods.
- Renamed the project to **Obsidian Image API Gateway** for better recognition.

---

## Pricing

This software is open-source and free to use. If you’d like to express your gratitude or support continued development, feel free to contribute using one of the options below:

- [![Paypal](https://img.shields.io/badge/paypal-haierkeys-yellow?style=social&logo=paypal)](https://paypal.me/haierkeys)
- [<img src="https://cdn.buymeacoffee.com/buttons/v2/default-yellow.png" alt="BuyMeACoffee" width="100">](https://www.buymeacoffee.com/haierkeys)
- <img src="https://raw.githubusercontent.com/haierkeys/obsidian-auto-image-remote-uploader/main/bmc_qr.png" style="width:120px;height:auto;">
- **Afdian:** [https://afdian.net/a/haierkeys](https://afdian.net/a/haierkeys)

---

# Getting Started

### Dockerized Installation

Suppose your server’s image storage path is set to `/data/storage/uploads`. Run the following commands sequentially:

```bash
# Pull the Docker image
docker pull haierkeys/obsidian-image-api-gateway:latest

# Create necessary directories
mkdir -p /data/image-api/config
mkdir -p /data/image-api/storage/logs
mkdir -p /data/image-api/storage/uploads

# Download the default configuration to the config directory
wget https://raw.githubusercontent.com/haierkeys/obsidian-image-api-gateway/main/configs/config.yaml -O /data/config/config.yaml

# Create and start the container
docker run -tid --name image-api \
        -p 8000:8000 -p 8001:8001 \
        -v /data/image-api/storage/logs/:/api/storage/logs/ \
        -v /data/image-api/storage/uploads/:/api/storage/uploads/ \
        -v /data/image-api/config/:/api/config/ \
        haierkeys/obsidian-image-api-gateway:latest
```

---

### Binary Installation

Download the latest release from [https://github.com/haierkeys/obsidian-image-api-gateway/releases](https://github.com/haierkeys/obsidian-image-api-gateway/releases).

Extract the files to a desired directory and execute the program.

---

### Configuration

The configuration file is located at `./configs/config.yaml`. Below is the default content:

```yaml
server:
  run-mode:
  # Server ports - Use `ip:port` (specific IP) or `:port` (listen on all IPs)
  http-port: :8000
  read-timeout: 60
  write-timeout: 60
  # Performance monitoring endpoint
  private-http-listen: :8001

security:
  # API authorization token for image uploads
  auth-token: 6666

app:
  default-page-size: 10
  max-page-size: 100
  default-context-timeout: 60
  log-save-path: storage/logs
  log-file: app.log

  temp-path: storage/temp
  # Prefix for API responses with uploaded image URLs
  upload-url-pre: https://image.diybeta.com
  # Upload size limit in MB
  upload-max-size: 5
  # Allowed image file types
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

# Local storage configuration
local-fs:
  enable: true
  # Enable built-in file URL access service
  httpfs-enable: true
  save-path: storage/uploads

# Aliyun OSS configuration
oss:
  enable: false
  custom-path: blog
  bucket-name:
  endpoint:
  access-key-id:
  access-key-secret:

# Cloudflare R2 configuration
cloudflu-r2:
  enable: true
  custom-path: blog
  bucket-name: image
  account-id:
  access-key-id:
  access-key-secret:

# Email error reporting
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

---

## TODO

---

## Other

**Obsidian Auto Image Remote Uploader**

[https://github.com/haierkeys/obsidian-auto-image-remote-uploader](https://github.com/haierkeys/obsidian-auto-image-remote-uploader)