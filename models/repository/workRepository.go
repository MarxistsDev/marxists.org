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
	err := repo.Db.Preload("Works").First(&work, id).Error
	fmt.Println("Work Title: ", work.Title)
	fmt.Println("Num of Articles: ", len(work.Works))
	return &work, err
}
