package comment

import (
	"context"
	"errors"
	"fmt"
)

var (
	ErrFetchingComment = errors.New("failed to fetch comment by id")
	ErrNotImplemented  = errors.New("not implemented")
)

// Comment - a representation of the comment structure
// for the service
type Comment struct {
	ID     string
	Slug   string
	Body   string
	Author string
}

// Store - Defines all our service that the service needs
// to operate
type Store interface {
	GetComment(context.Context, string) (Comment, error)
	PostComment(context.Context, Comment) (Comment, error)
	DeleteComment(context.Context, string) error
	UpdateComment(context.Context, string, Comment) (Comment, error)
}

// Service - contains all the business logic
type Service struct {
	Store Store
}

// NewService - returns a pointer to a new service
func NewService(store Store) *Service {
	return &Service{
		Store: store,
	}
}

func (s *Service) GetComment(ctx context.Context, id string) (Comment, error) {
	fmt.Println("Retrieving a comment")
	comment, err := s.Store.GetComment(ctx, id)
	if err != nil {
		fmt.Println(err)
		return Comment{}, ErrFetchingComment
	}
	return comment, nil
}

func (s *Service) UpdateComment(ctx context.Context, id string, updatedComment Comment) (Comment, error) {
	comment, err := s.Store.UpdateComment(ctx, id, updatedComment)
	if err != nil {
		fmt.Println("error updating comment")
		return Comment{}, nil
	}
	return comment, nil
}

func (s *Service) DeleteComment(ctx context.Context, id string) error {
	return s.Store.DeleteComment(ctx, id)
}

func (s *Service) PostComment(ctx context.Context, comment Comment) (Comment, error) {
	insertedComment, err := s.Store.PostComment(ctx, comment)
	if err != nil {
		return Comment{}, err
	}

	return insertedComment, nil
}
