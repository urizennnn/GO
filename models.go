package main

import (
	"time"

	"github.com/google/uuid"
	"github.com/urizennnn/GO-PROJECTS/internal/database"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	ApiKey    string    `json:"api_key"`
}

func databaseUserToUser(dbUser database.User) User {
	return User{
		ID:        dbUser.ID,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
		Name:      dbUser.Name,
		ApiKey:    dbUser.ApiKey,
	}
}

type Feed struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	Url       string    `json:"url"`
	UserId    uuid.UUID
}

func databasetoFeed(feed database.Feed) Feed {
	return Feed{
		ID:        feed.ID,
		CreatedAt: feed.CreatedAt,
		UpdatedAt: feed.UpdatedAt,
		Name:      feed.Name,
		Url:       feed.Url,
		UserId:    feed.UserID,
	}
}
func singleFeed(feed []database.Feed) []Feed {
	feeds := []Feed{}
	for _, feed := range feed {
		feeds = append(feeds, databasetoFeed(feed))
	}

	return feeds
}

type FollowsFeed struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	FeedID    uuid.UUID `json:"feed_id"`
	UserId    uuid.UUID
}

func databaseFollowsFeed (followFeed database.FollowsFeed) FollowsFeed{
	return FollowsFeed{
		ID: followFeed.FeedID,
		CreatedAt: followFeed.CreatedAt,
		UpdatedAt: followFeed.UpdatedAt,
		FeedID: followFeed.FeedID,
		UserId: followFeed.UserID,
	}
}

func UserFeed(feedFollow []database.FollowsFeed) []FollowsFeed {
	userFeed := []FollowsFeed{}
	for _, feedFollows := range feedFollow {
		userFeed = append(userFeed, databaseFollowsFeed(feedFollows))
	}

	return userFeed
}