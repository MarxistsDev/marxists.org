package models

// Author represents the Author table.
type Author struct {
	AuthorID int    `gorm:"primaryKey" json:"author_id"`
	Name     string `json:"name" binding:"required"`
	OldWorks string `json:"old_works"`
}

// TableName specifies the table name for the Author model
func (Author) TableName() string {
	return "Author"
}

// Glossary represents the Glossary table.
type Glossary struct {
	GlossaryID  int    `gorm:"primaryKey" json:"glossary_id"`
	AuthorID    int    `json:"author_id" binding:"required"`
	Name        string `json:"name" binding:"required"`
	Shortname   string `json:"shortname"`
	Image       string `json:"image"`
	Description string `json:"description"`
}

// Work represents the Work table.
type Work struct {
	WorkID          int    `gorm:"primaryKey" json:"work_id"`
	Title           string `json:"title"`
	Written         string `json:"written"`
	PublicationDate string `json:"publication_date"`
	Source          string `json:"source"`
	Translated      string `json:"translated"`
	Transcription   string `json:"transcription"`
	Copyright       string `json:"copyright"`
	OldWorksIndex   string `json:"old_works_index"`
}

// AuthorWork represents the Author_Work table for the Many-to-Many relationship between Author and Work.
type AuthorWork struct {
	AuthorWorkID int `gorm:"primaryKey" json:"author_work_id"`
	AuthorID     int `json:"author_id" binding:"required"`
	WorkID       int `json:"work_id" binding:"required"`
}

// Article represents the Article table.
type Article struct {
	ArticleID int    `gorm:"primaryKey" json:"article_id"`
	WorkID    int    `json:"work_id" binding:"required"`
	Title     string `json:"title" binding:"required"`
	Content   string `json:"content" binding:"required"`
	Note      string `json:"note"`
	OldWorks  string `json:"old_works"`
}

// Movement represents the Movement table.
type Movement struct {
	MovementID  int    `gorm:"primaryKey" json:"movement_id"`
	Name        string `json:"name" binding:"required"`
	OldMovement string `json:"old_movement"`
}

// Collection represents the Collection table.
type Collection struct {
	CollectionID  int    `json:"collection_id"`
	Name          string `json:"name" binding:"required"`
	OldCollection string `json:"old_collection"`
}

// MovementAuthor represents the Movement_Author table for the Many-to-Many relationship between Movement and Author.
type MovementAuthor struct {
	MovementAuthorID int `gorm:"primaryKey" json:"movement_author_id"`
	MovementID       int `json:"movement_id" binding:"required"`
	AuthorID         int `json:"author_id" binding:"required"`
}

// CollectionWork represents the Collection_Work table for the Many-to-Many relationship between Collection and Work.
type CollectionWork struct {
	CollectionWorkID int `gorm:"primaryKey" json:"collection_work_id"`
	CollectionID     int `json:"collection_id" binding:"required"`
	WorkID           int `json:"work_id" binding:"required"`
}
