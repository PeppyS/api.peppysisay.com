package comments

import (
	"context"
	"encoding/json"
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

		comment := Comment{ID: doc.Ref.ID}
		json.Unmarshal(bytes, &comment)

		return comment
	}).([]Comment), err
}

func (cs *CommentService) GetByID(ctx context.Context, id string) (Comment, error) {
	doc, err := cs.db.Collection("comments").Doc(id).Get(ctx)
	if err != nil {
		return Comment{}, err
	}

	bytes, err := json.Marshal(doc.Data())
	if err != nil {
		return Comment{}, err
	}

	comment := Comment{ID: doc.Ref.ID}
	json.Unmarshal(bytes, &comment)

	return comment, nil
}

func (cs *CommentService) New(ctx context.Context, postID string, comment string, name string) (string, error) {
	newComment := map[string]interface{}{
		// TODO re-use the current Comment struct
		"created_at":  time.Now(),
		"first_name":  name,
		"last_name":   name,
		"likes_count": 0,
		"post_id":     cs.db.Collection("posts").Doc(postID),
		"text":        comment,
	}

	docRef, _, err := cs.db.Collection("comments").Add(ctx, newComment)
	if err != nil {
		return "", err
	}

	return docRef.ID, nil
}

func (cs *CommentService) DeleteByID(ctx context.Context, commentID string) error {
	_, err := cs.db.Collection("comments").Doc(commentID).Delete(ctx)

	return err
}
