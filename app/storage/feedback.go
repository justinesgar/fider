package storage

import "github.com/getfider/fider/app/models"

// Idea contains read and write operations for ideas
type Idea interface {
	GetByID(ideaID int) (*models.Idea, error)
	GetByNumber(number int) (*models.Idea, error)
	GetCommentsByIdeaID(ideaID int) ([]*models.Comment, error)
	GetAll() ([]*models.Idea, error)
	Save(userID int, title, description string) (*models.Idea, error)
	AddComment(userID, ideaID int, content string) (int, error)
	AddSupporter(userID, ideaID int) error
	RemoveSupporter(userID, ideaID int) error
	SetResponse(ideaID int, text string, userID, status int) error
}
