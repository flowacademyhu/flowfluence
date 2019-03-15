package models

import (
	"context"

	"github.com/olivere/elastic"
	"github.com/rs/xid"
	"gopkg.in/couchbase/gocb.v1"
)

const (
	postIndexName = "posts"
)

// ChangeHandler contain bucket
type ChangeHandler struct {
	bucket *gocb.Bucket
}

// NewChangeHandler for init ChangeHandler
func NewChangeHandler(bucket *gocb.Bucket) *ChangeHandler {
	return &ChangeHandler{
		bucket: bucket,
	}
}

// Create user from user struct
func (handler *ChangeHandler) Create(change Change) (Change, error) {

	change.Hash = xid.New().String()

	_, error := handler.bucket.Upsert(change.Hash, change, 0)

	if error != nil {
		return Change{}, error
	}

	return change, nil

}

// PostHandler contain bucket
type PostHandler struct {
	client *elastic.Client
}

// NewPostHandler for init ChangeHandler
func NewPostHandler(client *elastic.Client) *PostHandler {
	return &PostHandler{
		client: client,
	}
}

// Create elastic Post document
func (handler *PostHandler) Create(ctx context.Context, post Post) (Post, error) {
	existIndex, err := handler.client.IndexExists(postIndexName).Do(ctx)
	if err != nil {
		return Post{}, err
	}

	if !existIndex {
		handler.client.CreateIndex(postIndexName).Do(ctx)
	}

	if post.ID == "" {
		post.ID = xid.New().String()
	}

	if _, err := handler.client.Index().Index(postIndexName).Type("post").BodyJson(post).Do(ctx); err != nil {
		return Post{}, err
	}

	return post, nil
}
