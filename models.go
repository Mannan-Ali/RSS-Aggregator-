package main

import (
	"time"

	"github.com/Mannan-Ali/RSS-Aggregator/internal/database"
	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	ApiKey    string    `json:"api_key"`
}

func databaseUsertoUser(dbUser database.User) User {
	return User{
		ID:        dbUser.ID,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
		Name:      dbUser.Name,
		ApiKey:    dbUser.ApiKey,
	}
}

//Read about txt point 15
// so in the createUserfunction we called in handler_user this function creates query and puts it into the database, then as we have returning*;
// in our query so we get returned with whatever is stored in user now then we
// pass this to databaseusertouser function where we assign it to a new struct user as we want same nameing convention as the
// database, as we cannot make changes to internal file, hence a public main user is returned.

// so basically now we have 2 user right??
// like one which has camelcasing and the main user that have and we are returning the main user but we do have normal one

type Feed struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	Url       string    `json:"url"`
	UserID    uuid.UUID `json:"user_id"`
}

func databaseFeedtoFeed(dbFeed database.Feed) Feed {
	return Feed{
		ID:        dbFeed.ID,
		CreatedAt: dbFeed.CreatedAt,
		UpdatedAt: dbFeed.UpdatedAt,
		Name:      dbFeed.Name,
		Url:       dbFeed.Url,
		UserID:    dbFeed.UserID,
	}
}

func databaseFeedstoFeeds(dbFeed []database.Feed) []Feed {
	feeds := []Feed{}
	for _, dbFeed := range dbFeed {
		feeds = append(feeds, databaseFeedtoFeed(dbFeed))
	}
	return feeds
}

type FeedFollower struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	UserID    uuid.UUID `json:"user_id"`
	FeedID    uuid.UUID `json:"feed_id"`
}

func databaseFeedFollowertoFeedFollower(dbFeedFollower database.FeedFollower) FeedFollower {
	return FeedFollower{
		ID:        dbFeedFollower.ID,
		CreatedAt: dbFeedFollower.CreatedAt,
		UpdatedAt: dbFeedFollower.UpdatedAt,
		UserID:    dbFeedFollower.UserID,
		FeedID:    dbFeedFollower.FeedID,
	}
}
func databaseUserFeedstoUserFeeds(dbUserfeeds []database.FeedFollower) []FeedFollower {
	userFeeds := []FeedFollower{}
	for _, dbUserFeed := range dbUserfeeds {
		userFeeds = append(userFeeds, databaseFeedFollowertoFeedFollower(dbUserFeed))
	}
	return userFeeds
}

type Post struct {
	ID          uuid.UUID `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Title       string    `json:"title"`
	Url         string    `json:"url"`
	Description *string   `json:"description"`
	PublishedAt time.Time `json:"published_at"`
	FeedID      uuid.UUID `json:"feed_id"`
}

func databasePostToPost(dbPost database.Post) Post {
	var description *string
	if dbPost.Description.Valid {
		description = &dbPost.Description.String
	}
	return Post{
		ID:          dbPost.ID,
		CreatedAt:   dbPost.CreatedAt,
		UpdatedAt:   dbPost.UpdatedAt,
		Title:       dbPost.Title,
		Url:         dbPost.Url,
		Description: description,
		PublishedAt: dbPost.PublishedAt,
		FeedID:      dbPost.FeedID,
	}
}

func databasePostsToPosts(dbPosts []database.Post) []Post {
	posts := []Post{}
	for _, dbPost := range dbPosts {
		posts = append(posts, databasePostToPost(dbPost))
	}
	return posts
}
