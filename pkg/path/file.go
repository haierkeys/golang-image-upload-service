package path

import (
    "os"
    "os/exec"
    "path/filepath"
    "strings"
)

func GetExePath() string {
    file, _ := exec.LookPath(os.Args[0])
    path, _ := filepath.Abs(file)
    index := strings.LastIndex(path, string(os.PathSeparator))
    return path[:index]
}
func Exists(path string) bool {
    _, err := os.Stat(path) // os.Stat获取文件信息
    if err != nil {
        if os.IsExist(err) {
            return true
        }
        return false
    }
    return true
}

func PathSuffixCheckAdd(path string, suffix string) string {
    if !strings.HasSuffix(path, suffix) {

        path = path + "/"
    }
    return path
}

func GetPath(path string, root string) string {
    path = strings.TrimPrefix(path, "/")
    realPath := ""
    if !Exists(root + path) {
        pwdDir, _ := os.Getwd()
        realPath = pwdDir + "/" + path
    } else {
        realPath = root + path
    }
    return realPath
}

// 判断所给路径是否为文件夹
func IsDir(path string) bool {
    s, err := os.Stat(path)
    if err != nil {
        return false
    }
    return s.IsDir()

}

// 判断所给路径是否为文件
func IsFile(path string) bool {
    return !IsDir(path)
}
