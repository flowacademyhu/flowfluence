package models

import (
	"context"
	"testing"
	"time"

	"github.com/olivere/elastic"
	"gopkg.in/couchbase/gocb.v1"
)

const couchbaseURL = "couchbase://localhost"

func TestChangeHandler(t *testing.T) {
	cluster, _ := gocb.Connect(couchbaseURL)
	cluster.Authenticate(gocb.PasswordAuthenticator{
		Username: "event_storage",
		Password: "event_storage",
	})

	bucket, error := cluster.OpenBucket("event_storage", "")

	if error != nil {
		t.Error("Error when opening bucket:", error)
		return
	}

	changeHandler := NewChangeHandler(bucket)

	if _, err := changeHandler.Create(
		Change{
			PartnerID:  "123",
			Event:      "CREATE",
			ModifiedBy: "Robi",
			Timestamp:  time.Now(),
			Post:       Post{}},
	); err != nil {
		t.Error("Error when change saving", err)
		return
	}

}

func TestPostHandler(t *testing.T) {
	ctx := context.Background()
	client, err := elastic.NewClient(
		elastic.SetURL("http://localhost:9200"),
		elastic.SetSniff(false))
	if err != nil {
		t.Error("Error when opening Elastic connections", err)
		return
	}

	postHandler := NewPostHandler(client)

	savedPost, err := postHandler.Create(ctx, Post{
		Title:       "New Post",
		AuthorID:    "1234",
		PartnerID:   "123",
		Permissions: []string{"ADMIN", "USER"},
		Sections:    []Section{},
	})

	if err != nil {
		t.Error("Error when post creating ", err)
		return
	}

	savedPost.Title = "Edited post"

	if _, err := postHandler.Update(ctx, savedPost); err != nil {
		t.Error("Error when post updating ", err)
		return
	}

	if _, err := postHandler.Get(ctx, savedPost.ID); err != nil {
		t.Error("Error when post get ", err)
		return
	}
}
