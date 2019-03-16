package models

import (
	"context"
	"encoding/json"

	"github.com/labstack/gommon/log"
	"github.com/olivere/elastic"
	"github.com/rs/xid"
	"gopkg.in/couchbase/gocb.v1"
)

const (
	postIndexName = "posts"
	postTypeName  = "post"
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
	handler.createPostIndexIfNotExists(ctx)

	if post.ID == "" {
		post.ID = xid.New().String()
	}

	if _, err := handler.client.Index().Index(postIndexName).Type("post").Id(post.ID).BodyJson(post).Do(ctx); err != nil {
		return Post{}, err
	}

	return post, nil
}

// Update elastic Post document
func (handler *PostHandler) Update(ctx context.Context, post Post) (Post, error) {
	handler.createPostIndexIfNotExists(ctx)

	if _, err := handler.client.Update().Index(postIndexName).Type(postTypeName).Id(post.ID).Doc(post).Do(ctx); err != nil {
		return Post{}, err
	}

	return post, nil
}

// Get elastic Post document by id
func (handler *PostHandler) Get(ctx context.Context, id string) (Post, error) {
	postResult, err := handler.client.Get().Index(postIndexName).Id(id).Do(ctx)

	if err != nil {
		return Post{}, err
	}

	return convertSearchResultToUsers(postResult)
}

func (handler *PostHandler) createPostIndexIfNotExists(ctx context.Context) {
	existIndex, err := handler.client.IndexExists(postIndexName).Do(ctx)
	if err != nil {
		log.Error("Error when index checking ", err)
		return
	}

	if !existIndex {
		if _, err := handler.client.CreateIndex(postIndexName).Do(ctx); err != nil {
			log.Error("Error when index creating ", err)
			return
		}
	}
}

func convertSearchResultToUsers(searchResult *elastic.GetResult) (Post, error) {
	if !searchResult.Found {
		return Post{}, nil
	}
	var post Post
	if err := json.Unmarshal(*searchResult.Source, &post); err != nil {
		return Post{}, err
	}
	return post, nil
}
