package posts

import (
	"context"
	"encoding/json"

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

	bytes, _ := json.Marshal(doc.Data())

	post := Post{ID: doc.Ref.ID}
	json.Unmarshal(bytes, &post)

	comments, _ := ps.commentService.GetAllByPostID(ctx, post.ID)
	post.Comments = comments

	return post, err
}
