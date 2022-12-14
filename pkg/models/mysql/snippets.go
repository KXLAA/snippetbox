package mysql

import (
	"database/sql"

	"github.com/KXLAA/snippetbox/pkg/models"
)

type SnippetModel struct {
	DB *sql.DB
}

func (model *SnippetModel) Insert(title, content, expires string) (int, error) {
	return 0, nil
}

func (model *SnippetModel) Get(id int) (*models.Snippet, error) {
	return nil, nil
}

func (model *SnippetModel) Latest() ([]*models.Snippet, error) {
	return nil, nil
}
