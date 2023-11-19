package repository

import (
	"gorm.io/gorm"
	"marxists.org/models"
)

type AuthorRepository struct {
	Db *gorm.DB
}

func (repo AuthorRepository) GetAll() []*models.Author {

	var authors []*models.Author
	repo.Db.Model(&models.Author{}).Find(&authors)
	return authors
}

func (repo AuthorRepository) Get(id int) (*models.Author, error) {

	var author models.Author
	err := repo.Db.Model(&models.Author{}).Joins("Glossary").Preload("Works").First(&author, id).Error
	return &author, err
}
