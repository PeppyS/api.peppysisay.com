package posts

import (
	"context"
	"encoding/json"
	"fmt"

	"cloud.google.com/go/firestore"
	"github.com/PeppyS/api.peppysisay.com/api/blog/comments"
	funk "github.com/thoas/go-funk"
)

type PostService struct {
	db             *firestore.Client
	commentService *comments.CommentService
}

type Post struct {
	ID       string             `json:"id"`
	Comments []comments.Comment `json:"comments"`
}

func NewService(db *firestore.Client, cs *comments.CommentService) *PostService {
	return &PostService{db, cs}
}

func (ps *PostService) GetAll(ctx context.Context) ([]Post, error) {
	docs, err := ps.db.Collection("posts").Documents(ctx).GetAll()
	if err != nil {
		return []Post{Post{}}, err
	}

	return funk.Map(docs, func(doc *firestore.DocumentSnapshot) Post {
		bytes, _ := json.Marshal(doc.Data())

		post := Post{ID: doc.Ref.ID}
		json.Unmarshal(bytes, &post)

		comments, _ := ps.commentService.GetAllByPostID(ctx, post.ID)
		post.Comments = comments

		return post
	}).([]Post), err
}

func (ps *PostService) GetByID(ctx context.Context, id string) (Post, error) {
	doc, err := ps.db.Collection("posts").Doc(id).Get(ctx)
	if err != nil {
		return Post{}, err
	}

	bytes, err := json.Marshal(doc.Data())
	if err != nil {
		return Post{}, err
	}

	post := Post{ID: doc.Ref.ID}
	json.Unmarshal(bytes, &post)

	comments, err := ps.commentService.GetAllByPostID(ctx, post.ID)
	if err != nil {
		return Post{}, err
	}

	post.Comments = comments

	return post, nil
}

func (ps *PostService) AddComment(ctx context.Context, postID string, comment string, name string) (string, error) {
	if postID == "" {
		return "", fmt.Errorf("Must provide the post ID")
	}

	if comment == "" {
		return "", fmt.Errorf("Must provide a comment")
	}

	if name == "" {
		return "", fmt.Errorf("Must provide a name")
	}

	_, err := ps.GetByID(ctx, postID)
	if err != nil {
		return "", fmt.Errorf("Invalid post ID given")
	}

	return ps.commentService.New(ctx, postID, comment, name)
}

func (ps *PostService) DeleteComment(ctx context.Context, postID string, commentID string) error {
	comment, err := ps.commentService.GetByID(ctx, commentID)
	if err != nil {
		return fmt.Errorf("Invalid comment ID")
	}

	if comment.Post.ID != postID {
		return fmt.Errorf("Comment does not belong to post")
	}

	return ps.commentService.DeleteByID(ctx, commentID)
}
