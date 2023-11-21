package repository

import (
	"gorm.io/gorm"
	"marxists.org/models"
)

type SearchRepository struct {
	Db *gorm.DB
}

func (repo SearchRepository) SearchArticle(query string) (*[]*models.Article, error) {

	var articles []*models.Article
	err := repo.Db.Where("search @@ websearch_to_tsquery(?)", query).Find(&articles).Error
	return &articles, err
}
