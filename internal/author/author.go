package author

// Author represents an author.
// @Summary Author struct to represent an author
// @Description Struct to represent an author
type Author struct {
	ID   int    `db:"id" json:"id" `
	Name string `db:"name" json:"name" example:"test_author"`
}
