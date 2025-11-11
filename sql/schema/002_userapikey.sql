-- +goose up
ALTER TABLE USERS
ADD COLUMN API_KEY VARCHAR(64) 
UNIQUE NOT NULL DEFAULT(
    encode(sha256(random()::text::bytea), 'hex')
);
--now as we have added values to the table and now we have delcared the api_key to be unique and not null that will give error as 
-- we already have 2 values with null values when atler table runs hence a default value 

-- +goose down 
ALTER TABLE USERS
DROP COLUMN API_KEY;

--the reason to make this migration is we want to  authenticate users based on thier 
--api keys now to do that we will have to store those api keys (about what type and kind of api ahead) so we alter table