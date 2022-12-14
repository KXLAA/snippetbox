package mysql

import (
	"database/sql"
	"errors"

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

	//SQL statement to execute
	stmt := `SELECT id, title, content, created, expires FROM snippets
			 WHERE expires > UTC_TIMESTAMP() AND id = ?`

	// Initialize a pointer to a new zeroed Snippet struct.
	snippet := &models.Snippet{}

	//execute the statement, since we are looking for a row we use the QueryRow() method
	row := model.DB.QueryRow(stmt, id)

	/// Use row.Scan() to copy the values from each field in sql.Row to the
	// corresponding field in the Snippet struct
	err := row.Scan(&snippet.ID, &snippet.Title, &snippet.Content, &snippet.Created, &snippet.Expires)

	if err != nil {
		// If the query returns no rows, then row.Scan() will return a sql.ErrNoRows error.
		// We use the errors.Is() function check for that error specifically, and return
		// our own models.ErrNoRecord error instead.
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrNoRecord
		} else {
			return nil, err
		}
	}

	//Return snippet
	return snippet, nil
}

func (model *SnippetModel) Latest() ([]*models.Snippet, error) {
	return nil, nil
}
