package path

import (
    "errors"
    "os"
    "os/exec"
    "path/filepath"
    "runtime"
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

// PathSuffixCheckAdd 检查路径后缀，如果没有则添加
func PathSuffixCheckAdd(path string, suffix string) string {
    if !strings.HasSuffix(path, suffix) {

        path = path + "/"
    }
    return path
}

func IsAbsPath(path string) bool {
    if runtime.GOOS == "windows" {
        // Windows系统
        if filepath.VolumeName(path) != "" {
            return true // 如果有盘符，则为绝对路径
        }
        return filepath.IsAbs(path) // 检查是否是绝对路径
    }
    // UNIX/Linux系统
    return filepath.IsAbs(path)
}

func GetAbsPath(path string, root string) (string, error) {

    // path = strings.TrimPrefix(path, "/")

    if root != "" {
        root = PathSuffixCheckAdd(root, "/")
    }

    realPath := root + path

    // 如果本身就是绝对路径 就直接返回
    if IsAbsPath(realPath) {

    } else {
        pwdDir, _ := os.Getwd()
        realPath = PathSuffixCheckAdd(pwdDir, "/") + path
    }
    if Exists(realPath) {
        return realPath, nil
    } else {
        return "", errors.New("path not exists")
    }
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
