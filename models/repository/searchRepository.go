package repository

import (
	"gorm.io/gorm"
	"marxists.org/models"
)

type SearchRepository struct {
	Db *gorm.DB
}

func (repo SearchRepository) SearchWork(query string) (*[]*models.Work, error) {
	var articles []*models.Work
	err := repo.Db.Raw("Select * from \"Work\" "+
		"WHERE search @@ websearch_to_tsquery('english', ?) or search @@ websearch_to_tsquery('simple', ?) "+
		"Order by ts_rank(search, websearch_to_tsquery('english',?)) + ts_rank(search, websearch_to_tsquery('simple',?)) DESC;",
		query, query, query, query).Find(&articles).Error
	return &articles, err
}

func (repo SearchRepository) SearchGlossary(query string) (*[]*models.Glossary, error) {
	var glossaries []*models.Glossary
	err := repo.Db.
		Raw("Select * from \"Glossary\" "+
			"WHERE author_id is not NULL and ((search @@ websearch_to_tsquery('english', ?) or search @@ websearch_to_tsquery('simple', ?)) or name ILIKE ?) "+
			"Order by ts_rank(search, websearch_to_tsquery('english',?)) + ts_rank(search, websearch_to_tsquery('simple',?)) DESC;",
			query, query, "%"+query+"%", query, query).
		Find(&glossaries).Error
	return &glossaries, err
}
