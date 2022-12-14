package mysql

import (
	"database/sql"

	"github.com/KXLAA/snippetbox/pkg/models"
)

type SnippetModel struct {
	DB *sql.DB
}

func (model *SnippetModel) Insert(title, content, expires string) (int, error) {
	//SQL statement to execute
	stmt := `INSERT INTO snippets (title, content, created, expires)
			VALUES(?, ?, UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY))`

	//execute the statement
	result, err := model.DB.Exec(stmt, title, content, expires)
	if err != nil {
		return 0, nil
	}

	//get the ID of newly inserted record in the table
	id, err := result.LastInsertId()
	if err != nil {
		return 0, nil
	}

	//Return id as an int type
	return int(id), nil
}

func (model *SnippetModel) Get(id int) (*models.Snippet, error) {
	return nil, nil
}

func (model *SnippetModel) Latest() ([]*models.Snippet, error) {
	return nil, nil
}
