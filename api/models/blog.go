package models

import (
	"gorm.io/gorm"
	"strings"
)

type Blog struct {
	gorm.Model
	Author ApiUser `gorm:"embedded"`
	Title  string           `gorm:"type:varchar(1000);not null" json:"title"`
	Body   string           `gorm:"size:1000;not null" json:"body"`
}
type ApiUser struct {
	UId   uint   `gorm:"primary key" json:"uid"`
	Email string `gorm:"type:varchar(100);unique_index" json:"email"`
	Name  string `gorm:"size:100;not null"              json:"name"`
}

func (b *Blog) Prepare() {
	b.Title = strings.TrimSpace(b.Title)
}

func (b *Blog) Save(db *gorm.DB) (*Blog, error) {
	var err error

	err = db.Debug().Create(&b).Error

	if err != nil {
		return &Blog{}, err
	}
	return b, nil
}

func (b *Blog) Get(id int, db *gorm.DB) (*Blog, error) {
	var err error
	blog := &Blog{}
	if err = db.Debug().First(blog, id).Error; err != nil {
		return nil, err
	}
	return blog, nil
}

func (b *Blog) GetAll(db *gorm.DB) (*[]Blog, error) {
	var blogs []Blog
	records := db.Debug().Order("created_at desc").Find(&blogs)
	if records.Error != nil {
		return &[]Blog{}, records.Error
	}
	return &blogs, nil
}

func (b *Blog) Put(id int, db *gorm.DB) (*Blog, error) {
	var err error
	err = db.Debug().Where("id = ?", id).Where(&Blog{Author: b.Author}).Updates(Blog{
		Title: b.Title,
		Body:  b.Body}).Error
	if err != nil {
		return &Blog{}, err
	}
	return b, nil
}

func (b *Blog) GetAllFromUser(db *gorm.DB) (*[]Blog,error){
	var blogs []Blog
	records := db.Debug().Where(&Blog{Author: b.Author}).Find(&blogs)
	if records.Error != nil{
		return &[]Blog{}, records.Error
	}
	return &blogs, nil
}