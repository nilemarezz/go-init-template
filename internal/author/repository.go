package author

import (
	"github.com/jmoiron/sqlx"
	"github.com/nilemarezz/go-init-template/pkg/logger"
)

type AuthorRepository interface {
	GetAllAuthors() ([]*Author, error)
	GetAuthorById(id int) (*Author, error)
	CreateAuthor(author *Author) error
	UpdateAuthor(author *Author, id int) error
}

type authorRepository struct {
	db *sqlx.DB
}

func NewAuthorRepository(db *sqlx.DB) AuthorRepository {
	return &authorRepository{db: db}
}

func (a authorRepository) GetAllAuthors() ([]*Author, error) {
	var authors []*Author
	logger.Info("query get all loggers")
	err := a.db.Select(&authors, "SELECT id, name FROM authors")
	return authors, err
}

func (a authorRepository) GetAuthorById(id int) (*Author, error) {
	var author Author
	err := a.db.Get(&author, "SELECT * FROM authors WHERE id = $1", id)
	return &author, err
}

func (a authorRepository) CreateAuthor(author *Author) error {
	// Insert the new author into the database
	_, err := a.db.Exec("INSERT INTO authors (name) VALUES ($1)", author.Name)
	if err != nil {
		return err
	}
	return nil
}

func (a authorRepository) UpdateAuthor(author *Author, id int) error {
	// Update the author in the database
	_, err := a.db.Exec("UPDATE authors SET name = $1 WHERE id = $2", author.Name, id)
	if err != nil {
		return err
	}
	return nil
}
