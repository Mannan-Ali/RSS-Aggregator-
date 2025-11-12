package main

import (
	"time"

	"github.com/Mannan-Ali/RSS-Aggregator/internal/database"
	"github.com/google/uuid"
)

// type User struct {
// 	ID        uuid.UUID
// 	CreatedAt time.Time
// 	UpdatedAt time.Time
// 	Name      string
// }

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
