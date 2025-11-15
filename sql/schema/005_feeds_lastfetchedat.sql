-- +goose up
ALTER TABLE FEEDS
ADD COLUMN LAST_FETCHED_AT TIMESTAMP;

-- +goose down 
ALTER TABLE FEEDS
DROP COLUMN LAST_FETCHED_AT;

-- we are adding this column so we can keep track of when we last 
---fetched our feed on our main website 