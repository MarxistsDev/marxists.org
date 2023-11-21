package repository

import (
	"fmt"

	"gorm.io/gorm"
	"marxists.org/models"
)

type WorkRepository struct {
	Db *gorm.DB
}

func (repo WorkRepository) Get(id int) (*models.Work, error) {

	var work models.Work
	err := repo.Db.Debug().Preload("Articles").First(&work, id).Error
	fmt.Println("Num of Articles: ", len(work.Articles))
	return &work, err
}
