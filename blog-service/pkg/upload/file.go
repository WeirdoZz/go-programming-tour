package upload

import (
	"blog-service/global"
	"blog-service/pkg/util"
	"io"
	"io/ioutil"
	"mime/multipart"
	"os"
	"path"
	"strings"
)

type FileType int

const TypeImage FileType = iota + 1

// GetFileName 获取MD5编码之后的文件名
func GetFileName(name string) string {
	ext := GetFileExt(name)
	// 从文件名中去掉后缀
	fileName := strings.TrimSuffix(name, ext)
	fileName = util.EncodeMD5(fileName)

	return fileName + ext
}

// GetFileExt 获取文件的后缀名
func GetFileExt(name string) string {
	return path.Ext(name)
}

// GetSavePath 获取文件保存路径
func GetSavePath() string {
	return global.AppSetting.UploadSavePath
}

// CheckSavePath 检查要存的路径是否存在
func CheckSavePath(dst string) bool {
	_, err := os.Stat(dst)
	return os.IsNotExist(err)
}

// CheckContainExt 查看该文件类型对应的后缀是否能够上传
func CheckContainExt(t FileType, name string) bool {
	ext := GetFileExt(name)
	ext = strings.ToUpper(ext)
	switch t {
	case TypeImage:
		for _, allowExt := range global.AppSetting.UploadImageAllowExts {
			if strings.ToUpper(allowExt) == ext {
				return true
			}
		}
	}

	return false
}

// CheckMaxSize 查看文件大小是否超过限制
func CheckMaxSize(t FileType, f multipart.File) bool {
	content, _ := ioutil.ReadAll(f)
	size := len(content)
	switch t {
	case TypeImage:
		if size >= global.AppSetting.UploadImageMaxSize*1024*1024 {
			return true
		}
	}
	return false
}

// CheckPermission 查看文件权限是否足够
func CheckPermission(dst string) bool {
	_, err := os.Stat(dst)
	return os.IsPermission(err)
}

// CreateSavePath 创建文件保存的路径
func CreateSavePath(dst string, perm os.FileMode) error {
	err := os.MkdirAll(dst, perm)
	if err != nil {
		return err
	}
	return nil
}

// SaveFile 将文件保存到dst
func SaveFile(file *multipart.FileHeader, dst string) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, src)
	return err
}
