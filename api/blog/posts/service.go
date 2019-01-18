package posts

import (
	"context"
	"encoding/json"
	"fmt"

	"cloud.google.com/go/firestore"
	"github.com/PeppyS/api.peppysisay.com/api/blog/comments"
	"github.com/gin-gonic/gin"
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

func (ps *PostService) AddComment(ctx *gin.Context, postID string, text string, name string) (comments.Comment, error) {
	if postID == "" {
		return comments.Comment{}, fmt.Errorf("Must provide the post ID")
	}

	if text == "" {
		return comments.Comment{}, fmt.Errorf("Must provide a comment")
	}

	if name == "" {
		return comments.Comment{}, fmt.Errorf("Must provide a name")
	}

	_, err := ps.GetByID(ctx, postID)
	if err != nil {
		return comments.Comment{}, fmt.Errorf("Invalid post ID given")
	}

	return ps.commentService.New(ctx, postID, text, name)
}

func (ps *PostService) DeleteComment(ctx *gin.Context, postID string, commentID string) error {
	comment, err := ps.commentService.GetByID(ctx, commentID)
	if err != nil {
		return fmt.Errorf("Invalid comment ID")
	}

	sessionID := ctx.Request.Header.Get("X-Session-ID")
	if comment.SessionID != sessionID {
		return fmt.Errorf("This comment does not belong to you")
	}

	if comment.Post.ID != postID {
		return fmt.Errorf("Comment does not belong to post")
	}

	return ps.commentService.DeleteByID(ctx, commentID)
}
