package file

import (
    "io"
    "os"
    "fmt"
    "log"
    "time"
    "path"
    "path/filepath"
    "errors"
    "strings"
)

// Flags 列表
const (
    // 只读模式
    O_RDONLY int = os.O_RDONLY
    // 只写模式
    O_WRONLY int = os.O_WRONLY
    // 可读可写
    O_RDWR   int = os.O_RDWR
    // 追加内容
    O_APPEND int = os.O_APPEND
    // 创建文件，如果文件不存在
    O_CREATE int = os.O_CREATE
    // 与创建文件一同使用，文件必须存在
    O_EXCL   int = os.O_EXCL
    // 打开一个同步的文件流
    O_SYNC   int = os.O_SYNC
    // 如果可能，打开时缩短文件
    O_TRUNC  int = os.O_TRUNC
)

// 将 path 中的 ‘/’ 转换为系统相关的路径分隔符
func FromSlash(s string) string {
    return filepath.FromSlash(s)
}

// 将 path 中平台相关的路径分隔符转换为 ‘/’
func ToSlash(s string) string {
    return filepath.ToSlash(s)
}

// 执行文件绝对路径
func SelfPath() string {
    pt, _ := filepath.Abs(os.Args[0])
    return pt
}

// 绝对路径
func RealPath(fp string) (string, error) {
    if path.IsAbs(fp) {
        return fp, nil
    }

    wd, err := os.Getwd()

    return filepath.Join(wd, fp), err
}

// 执行文件爱你目录
func SelfDir() string {
    return filepath.Dir(SelfPath())
}

// 文件目录名称
func Basename(fp string) string {
    return filepath.Base(fp)
}

// 获取文件夹名称
func Dir(fp string) string {
    return filepath.Dir(fp)
}

// 创建文件夹
func InsureDir(fp string) error {
    if IsExist(fp) {
        return nil
    }

    return os.MkdirAll(fp, os.ModePerm)
}

// 创建文件夹
func EnsureDir(fp string) error {
    return os.MkdirAll(fp, os.ModePerm)
}

// 是否可读
func IsReadable(file string) error {
    _, err := os.ReadFile(file)
    if err != nil {
        return err
    }

    return nil
}

// 创建文件夹并确认能够读写
func EnsureDirRW(dataDir string) error {
    err := EnsureDir(dataDir)
    if err != nil {
        return err
    }

    checkFile := fmt.Sprintf("%s/rw.%d", dataDir, time.Now().UnixNano())
    fd, err := Create(checkFile)
    if err != nil {
        if os.IsPermission(err) {
            return fmt.Errorf("open %s: rw permission denied", dataDir)
        }

        return err
    }

    if err := Close(fd); err != nil {
        return fmt.Errorf("close error: %s", err)
    }

    if err := Remove(checkFile); err != nil {
        return fmt.Errorf("remove error: %s", err)
    }

    return nil
}

// 创建
func Create(name string) (*os.File, error) {
    return os.Create(name)
}

// 删除
func Remove(name string) error {
    return os.Remove(name)
}

// 关闭
func Close(fd *os.File) error {
    return fd.Close()
}

// 后缀
func Ext(fp string) string {
    return filepath.Ext(fp)
}

// 重命名文件
func Rename(src string, target string) error {
    return os.Rename(src, target)
}

// 删除文件
func Unlink(fp string) error {
    return os.Remove(fp)
}

// 是否为文件
func IsFile(fp string) bool {
    f, e := os.Stat(fp)
    if e != nil {
        return false
    }

    return !f.IsDir()
}

// 判断所给路径是否为文件夹
func IsDir(fp string) bool {
    s, err := os.Stat(fp)
    if err != nil {
        return false
    }

    return s.IsDir()
}

// 文件是否存在
func IsExist(fp string) bool {
    _, err := os.Stat(fp)

    return err == nil || os.IsExist(err)
}

// 在目录里搜索文件
func SearchFile(filename string, paths ...string) (fullPath string, err error) {
    for _, pt := range paths {
        if fullPath = filepath.Join(pt, filename); IsExist(fullPath) {
            return
        }
    }

    err = fmt.Errorf("%s not found in paths", fullPath)
    return
}

// 文件最后更新时间
func FileMTime(fp string) (int64, error) {
    f, e := os.Stat(fp)
    if e != nil {
        return 0, e
    }

    return f.ModTime().Unix(), nil
}

// 文件大小，单位： bytes
func FileSize(fp string) (int64, error) {
    f, e := os.Stat(fp)
    if e != nil {
        return 0, e
    }

    return f.Size(), nil
}

// 列出文件夹
func DirsUnder(dirPath string) ([]string, error) {
    if !IsExist(dirPath) {
        return []string{}, nil
    }

    fs, err := os.ReadDir(dirPath)
    if err != nil {
        return []string{}, err
    }

    sz := len(fs)
    if sz == 0 {
        return []string{}, nil
    }

    ret := make([]string, 0, sz)
    for i := 0; i < sz; i++ {
        if fs[i].IsDir() {
            name := fs[i].Name()
            if name != "." && name != ".." {
                ret = append(ret, name)
            }
        }
    }

    return ret, nil
}

// 列出文件
func FilesUnder(dirPath string) ([]string, error) {
    if !IsExist(dirPath) {
        return []string{}, nil
    }

    fs, err := os.ReadDir(dirPath)
    if err != nil {
        return []string{}, err
    }

    sz := len(fs)
    if sz == 0 {
        return []string{}, nil
    }

    ret := make([]string, 0, sz)
    for i := 0; i < sz; i++ {
        if !fs[i].IsDir() {
            ret = append(ret, fs[i].Name())
        }
    }

    return ret, nil
}

// 打开文件
func MustOpenFile(fp string) *os.File {
    if strings.Contains(fp, "/") || strings.Contains(fp, "\\") {
        dir := Dir(fp)
        err := EnsureDir(dir)
        if err != nil {
            log.Fatalf("mkdir -p %s occur error %v", dir, err)
        }
    }

    f, err := os.OpenFile(fp, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
    if err != nil {
        log.Fatalf("open %s occur error %v", fp, err)
    }

    return f
}

/**
 * 拷贝文件夹,同时拷贝文件夹中的文件
 ×
 * @param srcPath 需要拷贝的文件夹路径
 * @param destPath 拷贝到的位置
 */
func CopyDir(srcPath string, destPath string) error {
    // 检测目录正确性
    if srcInfo, err := os.Stat(srcPath); err != nil {
        return err
    } else {
        if !srcInfo.IsDir() {
            e := errors.New("原始目录不是一个正确的目录！")
            return e
        }
    }

    if destInfo, err := os.Stat(destPath); err != nil {
        return err
    } else {
        if !destInfo.IsDir() {
            e := errors.New("目标目录不是一个正确的目录！")
            return e
        }
    }

    // 统一路径
    srcPath, _ = filepath.Abs(srcPath)
    destPath, _ = filepath.Abs(destPath)

    err := filepath.Walk(srcPath, func(path string, f os.FileInfo, err error) error {
        if f == nil {
            return err
        }

        if !f.IsDir() {
            // 重设为新路径
            destNewPath := strings.Replace(path, srcPath, destPath, -1)

            CopyFile(path, destNewPath)
        }

        return nil
    })

    return err
}

// 生成目录并拷贝文件
func CopyFile(src, dest string) (w int64, err error) {
    srcFile, err := os.Open(src)

    if err != nil {
        return
    }
    defer srcFile.Close()

    // 文件目录
    destPath, _ := filepath.Split(dest)

    // 目录不存在时
    if ok, _ := PathExists(destPath); !ok {
        // 创建目录
        err = os.MkdirAll(destPath, os.ModePerm)
        if err != nil {
            w = 0
            return
        }
    }

    dstFile, err := os.Create(dest)
    if err != nil {
        w = 0
        return
    }
    defer dstFile.Close()

    return io.Copy(dstFile, srcFile)
}

// 检测文件夹路径时候存在
func PathExists(path string) (bool, error) {
    _, err := os.Stat(path)
    if err == nil {
        return true, nil
    }

    if os.IsNotExist(err) {
        return false, nil
    }

    return false, err
}

// 写入文件
func WriteFile(filename string, contents string, flag ...int) (int, error) {
    // os.O_CREATE|os.O_RDWR|os.O_APPEND
    newFlag := os.O_CREATE|os.O_WRONLY
    if len(flag) > 0 {
        newFlag = flag[0]
    }

    data := []byte(contents)

    // 创建文件夹
    InsureDir(Dir(filename))

    fl, err := os.OpenFile(filename, newFlag, 0666)
    if err != nil {
        return 0, err
    }

    defer fl.Close()

    return fl.Write(data)
}

// 读取文件
func ReadFile(filename string) (string, error) {
    file, err := os.Open(filename)
    if err != nil {
        return "", err
    }
    defer file.Close()

    data, err2 := io.ReadAll(file)
    if err2 != nil {
        return "", err2
    }

    return string(data), nil
}

// 格式化数据大小
func FormatBytes(size int64) string {
    units := []string{" B", " KB", " MB", " GB", " TB", " PB"}

    s := float64(size)

    i := 0
    for ; s >= 1024 && i < 5; i++ {
        s /= 1024
    }

    return fmt.Sprintf("%.2f%s", s, units[i])
}

// 创建软链接
func Symlink(target, link string) error {
    return os.Symlink(target, link)
}

// 读取链接
func Readlink(link string) (string, error) {
    return os.Readlink(link)
}

// 是否为软链接
func IsSymlink(m os.FileMode) bool {
    return m&os.ModeSymlink != 0
}
