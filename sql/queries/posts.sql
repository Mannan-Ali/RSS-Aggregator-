-- name: CreatePost :one
INSERT INTO POSTS(ID,CREATED_AT,UPDATED_AT,
TITLE,URL,DESCRIPTION,PUBLISHED_AT,FEED_ID)
VALUES ($1,$2,$3,$4,$5,$6,$7,$8)
RETURNING *;


-- name: GetNewPostForUser :many
SELECT POSTS.* FROM POSTS 
JOIN FEED_FOLLOWER ON POSTS.FEED_ID = FEED_FOLLOWER.FEED_ID
WHERE FEED_FOLLOWER.USER_ID = $1
ORDER BY POSTS.PUBLISHED_AT DESC
LIMIT $2;


-- WHERE FEED_FOLLOWER.USER_ID = $1 filtering post for a specific user

--feed followr is following a feed,and eac feed belongs to a feed, so we can filter all the posts user is actually following
-- now we want the user to see the lastest posts alaways so we always need to fetch the latest post