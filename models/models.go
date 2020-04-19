package models

import (
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type FeedUrl struct {
	gorm.Model
	FeedUrl           string
	LastKickoffDate   *time.Time
	LastCompletedDate *time.Time
}

type FeedPull struct {
	gorm.Model
	FeedUrlID int
	FeedUrl   FeedUrl //`gorm:"foreignkey:FeedUrlID"`
}

type FeedItem struct {
	gorm.Model
	FeedUrlID       int `gorm:"unique_index:idx_feedid_guid"`
	FeedUrl         FeedUrl
	Title           string
	Description     string
	Content         string
	Link            string
	Updated         string
	UpdatedParsed   *time.Time
	Published       string
	PublishedParsed *time.Time
	GUID            string `gorm:"unique_index:idx_feedid_guid"` //`gorm:"unique;not null"`
	// Author          *Person
	// Image           *Image                   `json:"image,omitempty"`
	// Categories      []string                 `json:"categories,omitempty"`
	// Enclosures      []*Enclosure             `json:"enclosures,omitempty"`
	// DublinCoreExt   *ext.DublinCoreExtension `json:"dcExt,omitempty"`
	// ITunesExt       *ext.ITunesItemExtension `json:"itunesExt,omitempty"`
	// Extensions      ext.Extensions           `json:"extensions,omitempty"`
	// Custom          map[string]string        `json:"custom,omitempty"
}
