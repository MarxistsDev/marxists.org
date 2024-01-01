package models

import (
	"html/template"
)

// Author entity
type Author struct {
	AuthorID  uint   `gorm:"primaryKey;column:author_id" json:"author_id"`
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
	GlossaryID uint   `gorm:"primaryKey"`
	AuthorID   uint   `gorm:"foreignKey:Author"`
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
	WorkID          uint   `gorm:"primaryKey"`
	ParentWorkID    *uint  `gorm:"foreignKey:ParentWorkID"`
	Title           string `gorm:"not null"`
	Written         string //`gorm:"type:date"`
	PublicationDate string //`gorm:"type:date"`
	Source          string
	Translated      string
	Transcription   string
	Copyright       string
	OldWork         string
	Html            template.HTML
	Works           []Work       `gorm:"foreignKey:ParentWorkID"`
	Authors         []Author     `gorm:"many2many:author_works"`
	Collections     []Collection `gorm:"many2many:collection_works"`
}

func (Work) TableName() string {
	return "Work"
}

// Movement entity
type Movement struct {
	MovementID  uint   `gorm:"primaryKey"`
	Name        string `gorm:"not null"`
	OldMovement string
	Authors     []Author `gorm:"many2many:movement_authors"`
}

func (Movement) TableName() string {
	return "Movement"
}

// Collection entity
type Collection struct {
	CollectionID  uint   `gorm:"primaryKey"`
	Name          string `gorm:"not null"`
	OldCollection string
	Works         []Work `gorm:"many2many:collection_works"`
}

func (Collection) TableName() string {
	return "Collection"
}
