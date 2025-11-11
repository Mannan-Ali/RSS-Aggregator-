-- name: CreateUser :one
INSERT INTO USERS (ID,CREATED_AT,UPDATED_AT,NAME,api_key)
VALUES ($1,$2,$3,$4,encode(sha256(random()::text::bytea), 'hex'))
RETURNING *;


--we can still pass $5 but then we will have to use the sha256 function when we call this function 
--instaed we directly did it here

--we need to change this file as we have added new column apikey 
--before
-- INSERT INTO USERS (ID,CREATED_AT,UPDATED_AT,NAME)
-- VALUES ($1,$2,$3,$4)
-- RETURNING *;


-- name: GetUserByAPIKey :one
SELECT * FROM USERS
WHERE API_KEY = $1 LIMIT 1;
