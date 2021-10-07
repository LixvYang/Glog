package model

import (
	"glog/utils/errmsg"

	"gorm.io/gorm"
)

type Category struct {
	ID   uint   `gorm:"primary_key;auto_increment" json:"id"`
	Name string `gorm:"type:varchar(20);not null" json:"name"`
}

//check cate if exist
func CheckCategory(name string) (code int) {
	var cate Category

	db.Select("id").Where("name = ?", name).First(&cate)
	if cate.ID > 0 {
		return errmsg.ERROR_CATENAME_USED
	}
	return errmsg.SUCCSE
}

//create cate
func CreateCate(cate *Category) (code int) {
	err := db.Create(&cate)
	if err != nil {
		return errmsg.ERROR //500
	}
	return errmsg.SUCCSE
}

//getcate info
func GetCateInfo(id int) (Category, int) {
	var cate Category
	db.Where("id = ?", id).First(&cate)
	return cate, errmsg.SUCCSE
}

//get cate list
func GetCate(pageSize, pageNum int) ([]Category, int64) {
	var cate []Category
	var total int64
	err := db.Find(&cate).Limit(pageSize).Offset(((pageNum) - 1) * pageSize).Error
	db.Model(&cate).Count(&total)
	if err != nil && err != gorm.ErrRecordNotFound {
		return cate, 0
	}
	return cate, total
}

//exit cate
func EditCate(id int, data *Category) int {
	var cate Category
	var maps = make(map[string]interface{})
	maps["name"] = data.Name

	err = db.Model(&cate).Where("id = ? ", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

//delete cate
func DeleteCate(id int) int {
	var cate Category

	err = db.Where("id = ? ", id).Delete(&cate).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}
