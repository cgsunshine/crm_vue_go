package example

import (
	"errors"
	"fmt"
	"mime/multipart"
	"os"
	"strconv"
	"strings"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/example"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/upload"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: Upload
//@description: 创建文件上传记录
//@param: file model.ExaFileUploadAndDownload
//@return: error

func (e *FileUploadAndDownloadService) Upload(file *example.ExaFileUploadAndDownload) error {
	return global.GVA_DB.Create(&file).Error
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: FindFile
//@description: 查询文件记录
//@param: id uint
//@return: model.ExaFileUploadAndDownload, error

func (e *FileUploadAndDownloadService) FindFile(id uint) (example.ExaFileUploadAndDownload, error) {
	var file example.ExaFileUploadAndDownload
	err := global.GVA_DB.Where("id = ?", id).First(&file).Error
	return file, err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: FindFile
//@description: 查询文件记录
//@param: id uint
//@return: model.ExaFileUploadAndDownload, error

func (e *FileUploadAndDownloadService) FindIdFile(id string) (example.ExaFileUploadAndDownload, error) {
	var file example.ExaFileUploadAndDownload
	err := global.GVA_DB.Where("id = ?", id).First(&file).Error
	return file, err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteFile
//@description: 删除文件记录
//@param: file model.ExaFileUploadAndDownload
//@return: err error

func (e *FileUploadAndDownloadService) DeleteFile(file example.ExaFileUploadAndDownload) (err error) {
	var fileFromDb example.ExaFileUploadAndDownload
	fileFromDb, err = e.FindFile(file.ID)
	if err != nil {
		return
	}
	oss := upload.NewOss()
	if err = oss.DeleteFile(fileFromDb.Key); err != nil {
		return errors.New("文件删除失败")
	}
	err = global.GVA_DB.Where("id = ?", file.ID).Unscoped().Delete(&file).Error
	return err
}

// EditFileName 编辑文件名或者备注
func (e *FileUploadAndDownloadService) EditFileName(file example.ExaFileUploadAndDownload) (err error) {
	var fileFromDb example.ExaFileUploadAndDownload
	return global.GVA_DB.Where("id = ?", file.ID).First(&fileFromDb).Update("name", file.Name).Error
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetFileRecordInfoList
//@description: 分页获取数据
//@param: info request.PageInfo
//@return: list interface{}, total int64, err error

func (e *FileUploadAndDownloadService) GetFileRecordInfoList(info request.PageInfo) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	keyword := info.Keyword
	db := global.GVA_DB.Model(&example.ExaFileUploadAndDownload{})
	var fileLists []example.ExaFileUploadAndDownload
	if len(keyword) > 0 {
		db = db.Where("name LIKE ?", "%"+keyword+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Order("updated_at desc").Find(&fileLists).Error
	return fileLists, total, err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UploadFile
//@description: 根据配置文件判断是文件上传到本地或者七牛云
//@param: header *multipart.FileHeader, noSave string
//@return: file model.ExaFileUploadAndDownload, err error

func (e *FileUploadAndDownloadService) UploadFile(header *multipart.FileHeader, noSave string) (file example.ExaFileUploadAndDownload, err error) {
	oss := upload.NewOss()
	filePath, key, uploadErr := oss.UploadFile(header)
	if uploadErr != nil {
		panic(uploadErr)
	}
	s := strings.Split(header.Filename, ".")
	f := example.ExaFileUploadAndDownload{
		Url:  filePath,
		Name: header.Filename,
		Tag:  s[len(s)-1],
		Key:  key,
	}
	if noSave == "0" {
		return f, e.Upload(&f)
	}
	return f, nil
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetFileRecordInfoList
//@description: 分通过ids 获取图片
//@param: info request.PageInfo
//@return: list interface{}, total int64, err error

func (e *FileUploadAndDownloadService) GetFileRecordInfoIdsList(ids string) (fileLists []example.ExaFileUploadAndDownload, total int64, err error) {

	db := global.GVA_DB.Model(&example.ExaFileUploadAndDownload{})
	//var fileLists []example.ExaFileUploadAndDownload
	//if len(keyword) > 0 {
	//	db = db.Where("name LIKE ?", "%"+keyword+"%")
	//}

	// 将字符串转换为整数切片
	var idList []uint
	for _, idStr := range strings.Split(ids, ",") {
		id, err := strconv.ParseUint(idStr, 10, 32)
		if err != nil {
			fmt.Println("Error parsing ID:", err)
			continue
		}
		idList = append(idList, uint(id))
	}
	db.Where("id in (?)", idList)
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Order("updated_at desc").Find(&fileLists).Error
	return fileLists, total, err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetFileRecordInfoIdsString
//@description: 分通过ids 获取图片
//@param: info request.PageInfo
//@return: list interface{}, total int64, err error

func (e *FileUploadAndDownloadService) GetFileRecordInfoIdsString(ids string) (img string, total int64, err error) {

	db := global.GVA_DB.Model(&example.ExaFileUploadAndDownload{})
	//var fileLists []example.ExaFileUploadAndDownload
	//if len(keyword) > 0 {
	//	db = db.Where("name LIKE ?", "%"+keyword+"%")
	//}

	// 将字符串转换为整数切片
	var idList []uint
	for _, idStr := range strings.Split(ids, ",") {
		id, err := strconv.ParseUint(idStr, 10, 32)
		if err != nil {
			fmt.Println("Error parsing ID:", err)
			continue
		}
		idList = append(idList, uint(id))
	}
	db.Where("id in (?)", idList)
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	fileLists := make([]example.ExaFileUploadAndDownload, 0)

	err = db.Order("updated_at desc").Debug().Find(&fileLists).Error
	result := make([]string, len(fileLists))
	for i, list := range fileLists {
		result[i] = list.Url
	}
	img = strings.Join(result, ",")
	return img, total, err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UploadFile
//@description: 根据配置文件判断是文件上传到本地或者七牛云
//@param: header *multipart.FileHeader, noSave string
//@return: file model.ExaFileUploadAndDownload, err error

func (c *FileUploadAndDownloadService) DownloadFile(id string) (file []byte, file_name string, err error) {
	//oss := upload.NewOss()
	//filePath, key, uploadErr := oss.UploadFile(header)
	//if uploadErr != nil {
	//	panic(uploadErr)
	//}
	fileInfo, err := c.FindIdFile(id)
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
