package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/crm"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/upload"
	"mime/multipart"
	"os"
	"strings"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UploadFile
//@description: 根据配置文件判断是文件上传到本地或者七牛云
//@param: header *multipart.FileHeader, noSave string
//@return: file model.ExaFileUploadAndDownload, err error

func (c *CrmContactFileService) UploadFile(header *multipart.FileHeader, noSave string, contactId, sort *int) (file crm.CrmContactFile, err error) {
	oss := upload.NewOss()
	filePath, key, uploadErr := oss.UploadFile(header)
	if uploadErr != nil {
		panic(uploadErr)
	}
	s := strings.Split(header.Filename, ".")
	f := crm.CrmContactFile{
		Url:       filePath,
		Name:      header.Filename,
		Tag:       s[len(s)-1],
		Key:       key,
		ContactId: contactId,
		Sort:      sort,
	}
	if noSave == "0" {
		return f, c.Upload(f)
	}
	return f, nil
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UploadFile
//@description: 根据配置文件判断是文件上传到本地或者七牛云
//@param: header *multipart.FileHeader, noSave string
//@return: file model.ExaFileUploadAndDownload, err error

func (c *CrmContactFileService) DownloadFile(id string) (file []byte, file_name string, err error) {
	//oss := upload.NewOss()
	//filePath, key, uploadErr := oss.UploadFile(header)
	//if uploadErr != nil {
	//	panic(uploadErr)
	//}
	fileInfo, err := c.GetCrmContactFile(id)
	file_name = fileInfo.Name
	if err != nil {
		return nil, "", err
	}
	// 打开文件
	file, err = os.ReadFile(fileInfo.Url)
	if err != nil {
		return nil, "", err
	}

	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: Upload
//@description: 创建合同文件上传记录
//@param: file model.ExaFileUploadAndDownload
//@return: error

func (e *CrmContactFileService) Upload(file crm.CrmContactFile) error {
	return global.GVA_DB.Create(&file).Error
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: Upload
//@description: 创建合同文件上传记录
//@param: file model.ExaFileUploadAndDownload
//@return: error

func (e *CrmContactFileService) Sort(id int, sort string) error {
	return global.GVA_DB.Where("id = ?", id).Update("sort", sort).Error
}
