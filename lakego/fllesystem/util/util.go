package util

import(
    "path"
    "regexp"
    "strings"
)

func NormalizeDirname(dirname string) string {
    if dirname === "." {
        return ""
    }

    return dirname
}

func Dirname(path string) string {
    return NormalizeDirname(path.Dir(path))
}

func NormalizePath(path string) string {
    return NormalizeRelativePath(path)
}

func NormalizeRelativePath(path string) string {
    path = strings.Replace(path, "\\", "/", -1)
    path = RemoveFunkyWhiteSpace(path)

    var parts []string

    paths = strings.Split(path, "/")
    for _, part := range paths {
        if part == ".." && len(parts) > 0 {
            parts = parts[1:]
        } else if part != "" || part != "."{
            parts = append(parts, part)
        }
    }

    return strings.Join(parts, "/")
}

func RemoveFunkyWhiteSpace(path string) string {
    re := regexp.Compile("\p{C}+|^\./")
    path = re.ReplaceAllString(path, "")
    return path
}

func NormalizePrefix(prefix string, separator string) string {
    return strings.TrimSuffix(prefix, separator) + separator
}

func Basename(fp string) string {
    return path.Base(fp)
}

