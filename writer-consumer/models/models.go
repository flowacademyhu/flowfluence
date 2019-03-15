package models

import "time"

// Post represents a posts actual state
type Post struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	AuthorID    string    `json:"author_id"`
	PartnerID   string    `json:"partner_id"`
	Permissions []string  `json:"permissions"`
	Sections    []Section `json:"sections"`
}

// Section represents a section of a post
type Section struct {
	Hash        string `json:"hash"`
	ContentType string `json:"content_type"`
	Content     string `json:"content"`
	BeforeHash  string `json:"before_hash"`
}

// User represents a user
type User struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Type     string `json:"type"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

// Change represents a post's change
type Change struct {
	Hash       string    `json:"hash"`
	PartnerID  string    `json:"partner_id"`
	Event      string    `json:"event"`
	ModifiedBy string    `json:"modified_by"`
	Timestamp  time.Time `json:"timestamp"`
	Post       Post      `json:"post"`
}
