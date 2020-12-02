package models

import (
	"github.com/Shubhaankar-sharma/golang-blog-api/users/user-models"
	"gorm.io/gorm"
	"strings"
)

type Blog struct {
	gorm.Model
	Author user_models.User `gorm:"embedded"`
	Title  string           `gorm:"type:varchar(1000);not null" json:"title"`
	Body   string           `gorm:"size:1000;not null" json:"body"`
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
	err = db.Debug().Where("id = ?", id).Updates(Blog{
		Title: b.Title,
		Body:  b.Body}).Error
	if err != nil {
		return &Blog{}, err
	}
	return b, nil
}
