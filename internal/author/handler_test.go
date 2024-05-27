package author

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/nilemarezz/go-init-template/internal/errs"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockAuthorService struct {
	mock.Mock
}

func (m *MockAuthorService) GetAllAuthors() ([]*Author, error) {
	args := m.Called()
	return args.Get(0).([]*Author), args.Error(1)
}

func (m *MockAuthorService) GetAuthorById(id int) (*Author, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*Author), args.Error(1)
}

func (m *MockAuthorService) CreateAuthor(author *Author) error {
	args := m.Called(author)
	return args.Error(0)
}

func (m *MockAuthorService) UpdateAuthor(author *Author, id int) error {
	args := m.Called(author, id)
	return args.Error(0)
}

func TestGetAllAuthor(t *testing.T) {
	// Arrange
	mockService := new(MockAuthorService)
	handler := NewAuthorHandler(mockService)
	router := gin.Default()
	router.GET("/authors", handler.GetAllAuthor)
	expectedAuthors := []*Author{
		{ID: 1, Name: "John Doe"},
		{ID: 2, Name: "Jane Smith"},
	}
	mockService.On("GetAllAuthors").Return(expectedAuthors, nil)
	req, _ := http.NewRequest("GET", "/authors", nil)
	w := httptest.NewRecorder()

	// Act
	router.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusOK, w.Code)
	var responseAuthors []*Author
	err := json.Unmarshal(w.Body.Bytes(), &responseAuthors)
	assert.NoError(t, err)
	assert.Equal(t, expectedAuthors, responseAuthors)
	mockService.AssertExpectations(t)
}

func TestGetAllAuthor_InternalServerError(t *testing.T) {
	// Arrange
	mockService := new(MockAuthorService)
	handler := NewAuthorHandler(mockService)
	router := gin.Default()
	router.GET("/authors", handler.GetAllAuthor)

	expectedError := errors.New("some error")
	mockService.On("GetAllAuthors").Return(nil, expectedError)

	req, _ := http.NewRequest("GET", "/authors", nil)
	w := httptest.NewRecorder()

	// Act
	router.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusInternalServerError, w.Code)
	mockService.AssertExpectations(t)
}

func TestGetAuthorByID_Success(t *testing.T) {
	// Arrange
	mockService := new(MockAuthorService)
	handler := NewAuthorHandler(mockService)
	router := gin.Default()
	router.GET("/authors/:id", handler.GetAuthorByID)

	expectedAuthor := &Author{ID: 1, Name: "John Doe"}
	mockService.On("GetAuthorById", 1).Return(expectedAuthor, nil)

	req, _ := http.NewRequest("GET", "/authors/1", nil)
	w := httptest.NewRecorder()

	// Act
	router.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, `{"id":1,"name":"John Doe"}`, w.Body.String())
}

func TestGetAuthorByID_NotFound(t *testing.T) {
	// Arrange
	mockService := new(MockAuthorService)
	handler := NewAuthorHandler(mockService)
	router := gin.Default()
	router.GET("/authors/:id", handler.GetAuthorByID)

	mockService.On("GetAuthorById", 1).Return(nil, errs.NewNotFoundError("Author"))

	req, _ := http.NewRequest("GET", "/authors/1", nil)
	w := httptest.NewRecorder()

	// Act
	router.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestGetAuthorByID_InvalidID(t *testing.T) {
	// Arrange
	mockService := new(MockAuthorService)
	handler := NewAuthorHandler(mockService)
	router := gin.Default()
	router.GET("/authors/:id", handler.GetAuthorByID)

	req, _ := http.NewRequest("GET", "/authors/invalid", nil)
	w := httptest.NewRecorder()

	// Act
	router.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, 500, w.Code)
}

func TestGetAuthorByID_InternalServerError(t *testing.T) {
	// Arrange
	mockService := new(MockAuthorService)
	handler := NewAuthorHandler(mockService)
	router := gin.Default()
	router.GET("/authors/:id", handler.GetAuthorByID)

	mockService.On("GetAuthorById", 1).Return(nil, errors.New("some error"))

	req, _ := http.NewRequest("GET", "/authors/1", nil)
	w := httptest.NewRecorder()

	// Act
	router.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusInternalServerError, w.Code)
}
