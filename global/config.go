package global

import (
    "fmt"
    "os"

    "github.com/pkg/errors"
    "gopkg.in/yaml.v3"

    "github.com/haierspi/golang-image-upload-service/pkg/path"
)

var (
    Config config
)

type config struct {
    File       string
    Server     server     `yaml:"server"`
    App        app        `yaml:"app"`
    Email      email      `yaml:"email"`
    Security   security   `yaml:"security"`
    LocalFS    localFS    `yaml:"local-fs"`
    OSS        oss        `yaml:"storage-oss"`
    CloudfluR2 cloudfluR2 `yaml:"cloudflu-r2"`
    AWSS3      awsS3      `yaml:"aws-s3"`
}

// Server is a struct that holds the server settings
type server struct {
    // RunMode is a string that holds the run mode of the server
    RunMode string `yaml:"run-mode"`
    // HttpPort is a string that holds the http port of the server
    HttpPort string `yaml:"http-port"`
    // ReadTimeout is a duration that holds the read timeout of the server
    ReadTimeout int `yaml:"read-timeout"`
    // WriteTimeout is a duration that holds the write timeout of the server
    WriteTimeout int `yaml:"write-timeout"`
    // PrivateHttpListen is a string that holds the private http listen address of the server
    PrivateHttpListen string `yaml:"private-http-listen"`
}

type security struct {
    AuthToken string `yaml:"auth-token"`
}

type app struct {
    // 默认页面大小
    DefaultPageSize int `yaml:"default-page-size"`
    // 最大页面大小
    MaxPageSize int `yaml:"max-page-size"`
    // 默认上下文超时时间
    DefaultContextTimeout int `yaml:"default-context-timeout"`
    // 日志保存路径
    LogSavePath string `yaml:"log-save-path"`
    // 日志文件名
    LogFile string `yaml:"log-file"`

    // 上传临时路径
    TempPath string `yaml:"temp-path"`
    // 上传服务器URL
    UploadUrlPre string `yaml:"upload-url-pre"`
    // 上传图片最大尺寸
    UploadMaxSize int `yaml:"upload-max-size"`
    // 上传图片允许的扩展名
    UploadAllowExts []string `yaml:"upload-allow-exts"`
}

// StorageLocal struct
type localFS struct {
    Enable       bool   `yaml:"enable"`
    HttpfsEnable bool   `yaml:"httpfs-enable"`
    SavePath     string `yaml:"save-path"`
}

// OSS struct
type oss struct {
    Enable          bool   `yaml:"enable"`
    Endpoint        string `yaml:"endpoint"`
    BucketName      string `yaml:"bucket-name"`
    AccessKeyID     string `yaml:"access-key-id"`
    AccessKeySecret string `yaml:"access-key-secret"`
    CustomPath      string `yaml:"custom-path"`
}

// AWS S3 struct
type awsS3 struct {
    Enable          bool   `yaml:"enable"`
    Region          string `yaml:"region"`
    BucketName      string `yaml:"bucket-name"`
    AccessKeyID     string `yaml:"access-key-id"`
    AccessKeySecret string `yaml:"access-key-secret"`
    CustomPath      string `yaml:"custom-path"`
}

type cloudfluR2 struct {
    Enable          bool   `yaml:"enable"`
    AccountId       string `yaml:"account-id"`
    BucketName      string `yaml:"bucket-name"`
    AccessKeyID     string `yaml:"access-key-id"`
    AccessKeySecret string `yaml:"access-key-secret"`
    CustomPath      string `yaml:"custom-path"`
}

type email struct {
    ErrorReportEnable bool     `yaml:"error-report-enable"`
    Host              string   `yaml:"host"`
    Port              int      `yaml:"port"`
    UserName          string   `yaml:"username"`
    Password          string   `yaml:"password"`
    IsSSL             bool     `yaml:"is-ssl"`
    From              string   `yaml:"from"`
    To                []string `yaml:"to"`
}

// ConfigLoad 初始化
func ConfigLoad(f string) error {

    realpath, err := path.GetAbsPath(f, "")
    if err != nil {
        return err
    }

    fmt.Println("Config Absolute Path: " + realpath)

    Config.File = f
    file, err := os.ReadFile(f)
    if err != nil {
        return errors.Wrap(err, "read config file failed")
    }

    err = yaml.Unmarshal(file, &Config)
    if err != nil {
        return errors.Wrap(err, "parse config file failed")
    }
    return nil

}
