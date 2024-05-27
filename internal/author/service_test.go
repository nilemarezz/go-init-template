package author

import (
	"database/sql"
	"errors"
	"os"
	"testing"

	"github.com/nilemarezz/go-init-template/pkg/logger"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mock repository

type MockAuthorRepository struct {
	mock.Mock
}

func (m *MockAuthorRepository) GetAllAuthors() ([]*Author, error) {
	args := m.Called()
	return args.Get(0).([]*Author), args.Error(1)
}

func (m *MockAuthorRepository) GetAuthorById(id int) (*Author, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*Author), args.Error(1)
}

func (m *MockAuthorRepository) CreateAuthor(author *Author) error {
	args := m.Called(author)
	return args.Error(0)
}

func (m *MockAuthorRepository) UpdateAuthor(author *Author, id int) error {
	args := m.Called(author, id)
	return args.Error(0)
}

// start test case

func TestMain(m *testing.M) {
	// Initialize logger for tests
	logger.InitTestLogger()

	// Run all tests
	code := m.Run()

	// Exit with the appropriate exit code
	os.Exit(code)
}

func TestGetAllAuthors(t *testing.T) {
	// Arrange
	mockRepo := new(MockAuthorRepository)
	authorSvc := NewAuthorService(mockRepo)

	expectedAuthors := []*Author{
		{ID: 1, Name: "John Doe"},
		{ID: 2, Name: "Jane Smith"},
	}

	mockRepo.On("GetAllAuthors").Return(expectedAuthors, nil)

	// Act
	authors, err := authorSvc.GetAllAuthors()

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, expectedAuthors, authors)
	mockRepo.AssertExpectations(t)
}

func TestGetAuthorById(t *testing.T) {
	// Arrange
	mockRepo := new(MockAuthorRepository)
	authorSvc := NewAuthorService(mockRepo)

	expectedAuthor := &Author{ID: 1, Name: "John Doe"}

	mockRepo.On("GetAuthorById", 1).Return(expectedAuthor, nil)

	// Act
	author, err := authorSvc.GetAuthorById(1)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, expectedAuthor, author)
	mockRepo.AssertExpectations(t)
}

func TestCreateAuthor(t *testing.T) {
	// Arrange
	mockRepo := new(MockAuthorRepository)
	authorSvc := NewAuthorService(mockRepo)

	author := &Author{ID: 1, Name: "John Doe"}

	mockRepo.On("CreateAuthor", author).Return(nil)

	// Act
	err := authorSvc.CreateAuthor(author)

	// Assert
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestUpdateAuthor(t *testing.T) {
	// Arrange
	mockRepo := new(MockAuthorRepository)
	authorSvc := NewAuthorService(mockRepo)

	author := &Author{ID: 1, Name: "John Doe"}

	mockRepo.On("GetAuthorById", 1).Return(author, nil)
	mockRepo.On("UpdateAuthor", author, 1).Return(nil)

	// Act
	err := authorSvc.UpdateAuthor(author, 1)

	// Assert
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestUpdateAuthor_NotFound(t *testing.T) {
	// Arrange
	mockRepo := new(MockAuthorRepository)
	authorSvc := NewAuthorService(mockRepo)

	author := Author{ID: 1, Name: "John Doe"}

	mockRepo.On("GetAuthorById", 1).Return(nil, sql.ErrNoRows)

	// // Act
	err := authorSvc.UpdateAuthor(&author, 1)

	// Assert
	assert.Error(t, err)
	assert.Equal(t, err.Error(), "Author not found")
	mockRepo.AssertExpectations(t)
}

func TestUpdateAuthor_RepoError(t *testing.T) {
	// Arrange
	mockRepo := new(MockAuthorRepository)
	authorSvc := NewAuthorService(mockRepo)

	author := &Author{ID: 1, Name: "John Doe"}

	mockRepo.On("GetAuthorById", 1).Return(author, nil)
	mockRepo.On("UpdateAuthor", author, 1).Return(errors.New("some error"))

	// Act
	err := authorSvc.UpdateAuthor(author, 1)

	// Assert
	assert.Error(t, err)
	assert.EqualError(t, err, "some error")
	mockRepo.AssertExpectations(t)
}

func TestGetAuthorById_NotFound(t *testing.T) {
	// Arrange
	mockRepo := new(MockAuthorRepository)
	authorSvc := NewAuthorService(mockRepo)

	// Mock the repository to return sql.ErrNoRows
	mockRepo.On("GetAuthorById", 1).Return(nil, sql.ErrNoRows)

	// Act
	_, err := authorSvc.GetAuthorById(1)

	// Assert
	assert.Equal(t, err.Error(), "Author not found")
	mockRepo.AssertExpectations(t)
}

func TestGetAuthorById_ReturnError(t *testing.T) {
	// Arrange
	mockRepo := new(MockAuthorRepository)
	authorSvc := NewAuthorService(mockRepo)

	expectedError := errors.New("some error")

	// Mock the repository to return an error other than sql.ErrNoRows
	mockRepo.On("GetAuthorById", 1).Return(nil, expectedError)

	// Act
	_, err := authorSvc.GetAuthorById(1)

	// Assert
	assert.Error(t, err)
	assert.Equal(t, expectedError, err)
	mockRepo.AssertExpectations(t)
}
