package repository

import (
	"fmt"

	"gorm.io/gorm"
	"marxists.org/models"
)

type AuthorRepository struct {
	Db *gorm.DB
}

func (repo AuthorRepository) GetAll() []*models.Author {

	var authors []*models.Author
	repo.Db.Find(&authors)
	return authors
}

func (repo AuthorRepository) Get(id int, page *int) (*models.Author, error) {

	var author models.Author
	err := repo.Db.Joins("Glossary").Preload("Works", func(tx *gorm.DB) *gorm.DB {
		return tx.Limit(25 * (*page + 1))
	}).First(&author, id).Error
	fmt.Println("Num of Works: ", len(author.Works))
	return &author, err
}
