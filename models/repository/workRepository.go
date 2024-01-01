package repository

import (
	"fmt"

	"gorm.io/gorm"
	"marxists.org/models"
)

type WorkRepository struct {
	Db *gorm.DB
}

func (repo WorkRepository) Get(id uint) (*models.Work, error) {
	var work models.Work
	err := repo.Db.Preload("Works").First(&work, id).Error
	fmt.Println("Work Title: ", work.Title)
	fmt.Println("Num of Articles: ", len(work.Works))

	fmt.Println("Parent: ", work.ParentWorkID)

	if work.ParentWorkID != nil {
		// This could easaly create a infite loop, and take down the server
		return repo.Get(*work.ParentWorkID)
	}

	return &work, err
}
