// handler.go
package author

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"

	"github.com/nilemarezz/go-init-template/internal/errs"
	httputil "github.com/nilemarezz/go-init-template/internal/util"
)

func SetupRouter(router *gin.Engine, db *sqlx.DB) {

	authorRepo := NewAuthorRepository(db)
	authorService := NewAuthorService(authorRepo)
	handler := NewAuthorHandler(authorService)

	authorRoutes := router.Group("/authors")
	{
		authorRoutes.GET("/", handler.GetAllAuthor)
		authorRoutes.GET("/:id", handler.GetAuthorByID)
		authorRoutes.POST("/", handler.CreateAuthor)
		authorRoutes.PUT("/", handler.UpdateAuthor)
		// Add other routes like GET, PUT, DELETE here
	}
}

type AuthorHandler struct {
	service AuthorService
}

func NewAuthorHandler(service AuthorService) *AuthorHandler {
	return &AuthorHandler{service: service}
}

// GetAllAuthor fetches all authors.
// @Summary Get all authors
// @Description Retrieve a list of all authors
// @Produce json
// @Success 200 {array} Author
// @Failure 500 {object} httputil.HTTPError
// @Router /authors [get]
func (h *AuthorHandler) GetAllAuthor(c *gin.Context) {
	authors, err := h.service.GetAllAuthors()
	if err != nil {
		httputil.NewError(c, http.StatusInternalServerError, err)
		return
	}
	c.JSON(200, authors)
}

// GetAuthorByID retrieves an author by ID.
// @Summary Get an author by ID
// @Description Retrieve an author by its ID
// @Produce json
// @Param id path int true "Author ID"
// @Success 200 {object} Author
// @Failure 400 {object} httputil.HTTPError "Invalid ID format"
// @Failure 404 {object} httputil.HTTPError "Author not found"
// @Failure 500 {object} httputil.HTTPError "Internal Server Error"
// @Router /authors/{id} [get]
func (h *AuthorHandler) GetAuthorByID(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		httputil.NewError(c, http.StatusInternalServerError, err)
		return
	}

	authors, err := h.service.GetAuthorById(id)

	if errHandle, ok := err.(*errs.NotFoundError); ok {
		httputil.NewError(c, http.StatusNotFound, errHandle)
		return
	} else if err != nil {
		httputil.NewError(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(200, authors)
}

// CreateAuthor creates a new author.
// @Summary Create a new author
// @Description Create a new author with the provided data
// @Accept json
// @Produce json
// @Param author body Author true "Author object"
// @Success 201
// @Failure 400 {object} httputil.HTTPError "Bad request"
// @Failure 500 {object} httputil.HTTPError "Internal Server Error"
// @Router /authors [post]
func (h *AuthorHandler) CreateAuthor(c *gin.Context) {
	var newAuthor Author
	if err := c.BindJSON(&newAuthor); err != nil {
		httputil.NewError(c, http.StatusInternalServerError, err)
		return
	}

	err := h.service.CreateAuthor(&newAuthor)
	if err != nil {
		httputil.NewError(c, http.StatusInternalServerError, err)
		return
	}

	c.Status(201)
}

// UpdateAuthor updates an existing author.
// @Summary Update an existing author
// @Description Update an existing author with the provided data
// @Accept json
// @Produce json
// @Param author body Author true "Author object"
// @Success 200
// @Failure 400 {object} httputil.HTTPError "Bad request"
// @Failure 404 {object} httputil.HTTPError "Author not found"
// @Failure 500 {object} httputil.HTTPError "Internal Server Error"
// @Router /authors [put]
func (h *AuthorHandler) UpdateAuthor(c *gin.Context) {
	var updatedAuthor Author
	if err := c.BindJSON(&updatedAuthor); err != nil {
		httputil.NewError(c, http.StatusInternalServerError, err)
		return
	}

	err := h.service.UpdateAuthor(&updatedAuthor, updatedAuthor.ID)

	if err != nil {
		httputil.NewError(c, http.StatusInternalServerError, err)
		return
	}

	c.Status(200)
}
