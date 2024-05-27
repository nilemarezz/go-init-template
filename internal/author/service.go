package author

import (
	"database/sql"

	"github.com/nilemarezz/go-init-template/internal/errs"
)

type AuthorService interface {
	GetAllAuthors() ([]*Author, error)
	GetAuthorById(id int) (*Author, error)
	CreateAuthor(author *Author) error
	UpdateAuthor(author *Author, id int) error
}

type authorService struct {
	repo AuthorRepository
}

func NewAuthorService(repo AuthorRepository) AuthorService {
	return &authorService{repo: repo}
}

func (a authorService) GetAllAuthors() ([]*Author, error) {
	return a.repo.GetAllAuthors()
}

func (a authorService) GetAuthorById(id int) (*Author, error) {
	author, err := a.repo.GetAuthorById(id)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("Author")
		}
		return nil, err
	}

	return author, nil
}

func (a authorService) CreateAuthor(author *Author) error {
	return a.repo.CreateAuthor(author)
}

func (a authorService) UpdateAuthor(author *Author, id int) error {
	// Check if author exists
	_, err := a.repo.GetAuthorById(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return errs.NewNotFoundError("Author")
		}
		return err
	}

	// Update author
	err = a.repo.UpdateAuthor(author, id)
	if err != nil {
		return err
	}

	return nil
}
