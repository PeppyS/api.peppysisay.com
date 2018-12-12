package comments

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"cloud.google.com/go/firestore"
	funk "github.com/thoas/go-funk"
)

type CommentService struct {
	db *firestore.Client
}

type Comment struct {
	ID         string                `json:"id"`
	Post       firestore.DocumentRef `json:"post_id"`
	Parent     firestore.DocumentRef `json:"parent_id,omitempty"`
	CreatedAt  time.Time             `json:"created_at"`
	FirstName  string                `json:"first_name"`
	LastName   string                `json:"last_name"`
	LikesCount int                   `json:"likes_count"`
	Text       string                `json:"text"`
}

func NewService(db *firestore.Client) *CommentService {
	return &CommentService{db}
}

func (c Comment) MarshalJSON() ([]byte, error) {
	type CommentAlias Comment

	return json.Marshal(struct {
		CommentAlias
		Post      string `json:"post_id"`
		Parent    string `json:"parent_id,omitempty"`
		CreatedAt string `json:"created_at"`
	}{
		CommentAlias: CommentAlias(c),
		Post:         c.Post.ID,
		Parent:       c.Parent.ID,
		CreatedAt:    c.CreatedAt.String(),
	})
}

func (cs *CommentService) GetAllByPostID(ctx context.Context, postID string) ([]Comment, error) {
	docs, err := cs.db.Collection("comments").
		Where("post_id", "==", cs.db.Collection("posts").Doc(postID)).
		Documents(ctx).GetAll()

	return funk.Map(docs, func(doc *firestore.DocumentSnapshot) Comment {
		bytes, _ := json.Marshal(doc.Data())
		fmt.Println(string(bytes))

		comment := Comment{ID: doc.Ref.ID}
		json.Unmarshal(bytes, &comment)

		return comment
	}).([]Comment), err
}

func (cs *CommentService) GetByID(id string) {

}
