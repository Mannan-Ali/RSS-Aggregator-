-- name: CreateFeedFollower :one
INSERT INTO FEED_FOLLOWER(ID,CREATED_AT,UPDATED_AT,USER_ID,FEED_ID)
VALUES ($1,$2,$3,$4,$5)
RETURNING *;



-- name: GetAllFollowersFeeds :many
SELECT * FROM FEED_FOLLOWER WHERE USER_ID = $1; 

-- name: UnfollowUserFeed :exec
DELETE FROM FEED_FOLLOWER WHERE ID = $1 AND USER_ID = $2; 