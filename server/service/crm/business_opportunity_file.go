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

func (c *CrmBusinessOpportunityFileService) UploadFile(header *multipart.FileHeader, noSave string, businessOpportunityId, sort *int) (file crm.CrmBusinessOpportunityFile, err error) {
	oss := upload.NewOss()
	filePath, key, uploadErr := oss.UploadFile(header)
	if uploadErr != nil {
		panic(uploadErr)
	}
	s := strings.Split(header.Filename, ".")
	f := crm.CrmBusinessOpportunityFile{
		Url:                   filePath,
		Name:                  header.Filename,
		Tag:                   s[len(s)-1],
		Key:                   key,
		BusinessOpportunityId: businessOpportunityId,
		Sort:                  sort,
	}
	if noSave == "0" {
		return f, c.Upload(f)
	}
	return f, nil
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: Upload
//@description: 创建商机文件上传记录
//@param: file model.ExaFileUploadAndDownload
//@return: error

func (e *CrmBusinessOpportunityFileService) Upload(file crm.CrmBusinessOpportunityFile) error {
	return global.GVA_DB.Create(&file).Error
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: Upload
//@description: 排序修改
//@param: file model.ExaFileUploadAndDownload
//@return: error

func (e *CrmBusinessOpportunityFileService) Sort(id int, sort string) error {
	return global.GVA_DB.Where("id = ?", id).Update("sort", sort).Error
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UploadFile
//@description: 根据配置文件判断是文件上传到本地或者七牛云
//@param: header *multipart.FileHeader, noSave string
//@return: file model.ExaFileUploadAndDownload, err error

func (c *CrmBusinessOpportunityFileService) DownloadFile(id string) (file []byte, file_name string, err error) {
	//oss := upload.NewOss()
	//filePath, key, uploadErr := oss.UploadFile(header)
	//if uploadErr != nil {
	//	panic(uploadErr)
	//}
	fileInfo, err := c.GetCrmBusinessOpportunityFile(id)
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
