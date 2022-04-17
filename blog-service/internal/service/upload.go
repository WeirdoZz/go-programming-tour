package service

import (
	"blog-service/global"
	"blog-service/pkg/upload"
	"errors"
	"mime/multipart"
	"os"
)

type FileInfo struct {
	Name      string
	AccessUrl string
}

// UploadFile 上传文件的具体服务
func (svc *Service) UploadFile(fileType upload.FileType, file multipart.File,
	fileHeader *multipart.FileHeader) (*FileInfo, error) {

	fileName := upload.GetFileName(fileHeader.Filename)
	uploadSavePath := upload.GetSavePath()
	dst := uploadSavePath + "/" + fileName
	if !upload.CheckContainExt(fileType, fileName) {
		return nil, errors.New("file suffix is not supported. ")
	}
	if upload.CheckSavePath(uploadSavePath) {
		err := upload.CreateSavePath(uploadSavePath, os.ModePerm)
		if err != nil {
			return nil, errors.New("failed to create save directory. ")
		}
	}
	if upload.CheckMaxSize(fileType, file) {
		return nil, errors.New("exceeded maximum file limit. ")
	}
	if upload.CheckPermission(uploadSavePath) {
		return nil, errors.New("insufficient file permission. ")
	}
	if err := upload.SaveFile(fileHeader, dst); err != nil {
		return nil, err
	}

	accessUrl := global.AppSetting.UploadServerUrl + "/" + fileName
	return &FileInfo{Name: fileName, AccessUrl: accessUrl}, nil
}
