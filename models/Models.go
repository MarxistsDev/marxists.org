package models

import (
	"html/template"
)

// Author entity
type Author struct {
	AuthorID  int    `gorm:"primaryKey;column:author_id" json:"author_id"`
	Name      string `gorm:"not null"`
	OldWorks  string
	Works     []*Work    `gorm:"many2many:author_works;joinTable:author_works;joinForeignKey:author_author_id;joinReferences:work_work_id"`
	Glossary  Glossary   `gorm:"foreignKey:AuthorID"`
	Movements []Movement `gorm:"many2many:movement_authors;foreignKey:AuthorID"`
	//Collections []Collection `gorm:"many2many:collection_authors"`
}

// TableName specifies the table name for the Author model
func (Author) TableName() string {
	return "Author"
}

// Glossary entity
type Glossary struct {
	GlossaryID int `gorm:"primaryKey"`
	AuthorID   int
	Name       string `gorm:"not null"`
	//ShortName   string
	Image       string
	Description template.HTML
}

func (Glossary) TableName() string {
	return "Glossary"
}

// Work entity
type Work struct {
	WorkID          int    `gorm:"primaryKey"`
	Title           string `gorm:"not null"`
	Written         string `gorm:"type:date"`
	PublicationDate string `gorm:"type:date"`
	Source          string
	Translated      string
	Transcription   string
	Copyright       string
	OldWorksIndex   string
	Authors         []Author     `gorm:"many2many:author_works"`
	Articles        []Article    `gorm:"foreignKey:WorkID"`
	Collections     []Collection `gorm:"many2many:collection_works"`
}

func (Work) TableName() string {
	return "Work"
}

/*type AuthorWork struct {
	//AuthorWorkID int `gorm:"primaryKey" json:"author_work_id"`
	AuthorID int `primaryKey;gorm:"column:author_author_id" json:"author_id" binding:"required"`
	WorkID   int `primaryKey;gorm:"column:work_work_id" json:"work_id" binding:"required"`
}

func (AuthorWork) TableName() string {
	return "AuthorWork"
}*/

// Article entity
type Article struct {
	ArticleID int    `gorm:"primaryKey"`
	WorkID    int    `gorm:"not null"`
	Title     string `gorm:"not null"`
	Content   string
	Note      string
	OldWorks  string
	//search    string `gorm:"type:tsvector"`
}

func (Article) TableName() string {
	return "Article"
}

// Movement entity
type Movement struct {
	MovementID  int    `gorm:"primaryKey"`
	Name        string `gorm:"not null"`
	OldMovement string
	Authors     []Author `gorm:"many2many:movement_authors"`
}

func (Movement) TableName() string {
	return "Movement"
}

// Collection entity
type Collection struct {
	CollectionID  int    `gorm:"primaryKey"`
	Name          string `gorm:"not null"`
	OldCollection string
	Works         []Work `gorm:"many2many:collection_works"`
}

func (Collection) TableName() string {
	return "Collection"
}

// MovementAuthor entity
/*type MovementAuthor struct {
	MovementAuthorID int `gorm:"primaryKey"`
	MovementID       int `gorm:"not null"`
	AuthorID         int `gorm:"not null"`
}

// CollectionWork entity
type CollectionWork struct {
	CollectionWorkID int `gorm:"primaryKey"`
	CollectionID     int `gorm:"not null"`
	WorkID           int `gorm:"not null"`
}*/
