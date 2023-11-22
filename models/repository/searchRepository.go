package repository

import (
	"fmt"
	"html/template"
	"slices"
	"strings"

	"gorm.io/gorm"
	"marxists.org/models"
)

type SearchRepository struct {
	Db *gorm.DB
}

func (repo SearchRepository) SearchArticle(query string) (*[]*models.Article, error) {

	var articles []*models.Article
	err := repo.Db.Raw("Select * from \"Article\" "+
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

// This has an offsert since the toLower function converts everything to UTF-8, but some of the text uses
// HTML4 encoding (ISO)
func HighlightMatchingWords(html template.HTML, query string) template.HTML {

	text := string(html)

	stopWords := []string{"and", "the", "is"}
	query = strings.ToLower(query)
	textLower := strings.ToLower(text)

	fmt.Println("Query", query)

	queryTokens := strings.Split(query, " ")

	for _, element := range queryTokens {
		fmt.Println("Elelment:", element)
		if !slices.Contains(stopWords, element) {
			startIndex := 0
			for {
				index := strings.Index(textLower[startIndex:], element)
				if index == -1 {
					break
				}
				startIndex += index
				endIndex := startIndex + len(element)

				fmt.Println("Word: ", textLower[startIndex:endIndex])

				if !slices.Contains(stopWords, text[startIndex:endIndex]) {
					text = text[:startIndex] + "<mark>" + text[startIndex:endIndex] + "</mark>" + text[endIndex:]
				}

				startIndex = endIndex
			}
		}
	}

	MAX_CONTENT_SIZE := 500

	if len(text) > MAX_CONTENT_SIZE {
		text = text[:MAX_CONTENT_SIZE]
	}
	return template.HTML(text)
}
