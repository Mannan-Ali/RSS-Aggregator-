-- name: CreateFeed :one
INSERT INTO FEEDS (ID,CREATED_AT,UPDATED_AT,NAME,URL,USER_ID)
VALUES ($1,$2,$3,$4,$5,$6)
RETURNING *;


-- name: GetFeeds :many
SELECT * FROM FEEDS; 

-- name: GetNextFeedsToFech :many
SELECT * FROM FEEDS 
ORDER BY LAST_FETCHED_AT ASC NULLS FIRST 
LIMIT $1;


-- ORDER BY LAST_FETCHED_AT ASC NULLS FIRST: This is the most important part.
-- NULLS FIRST: It prioritizes any feeds that have NULL in their LAST_FETCHED_AT column. This means any brand-new feed that has never been fetched will always be picked first.
-- ASC (Ascending): If all feeds have been fetched at least once, it then sorts them by the LAST_FETCHED_AT date, from oldest to newest and then the oldest feeds get selected to be fetched
-- LIMIT 1: After sorting, it only takes the single feed at the top of the list.

-- name: MarkFeedAsFetched :one
UPDATE FEEDS 
SET LAST_FETCHED_AT = NOW(),
UPDATED_AT = NOW()
WHERE ID = $1
RETURNING *;

--  BOTH THE FUNCTION RUNS TOGETHER AS ONE
-- It calls getNextFeedToFech to get the highest-priority feed.
-- It fetches the content from that feed's URL.
-- If the fetch is successful, it calls MarkFeedAsFetched with the feed's ID to update its LAST_FETCHED_AT timestamp.
-- It waits for a short period (e.g., a minute) and then repeats the process.
-- This creates a continuous cycle that ensures all feeds are regularly updated, with new feeds and older ones getting priority