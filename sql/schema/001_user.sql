-- +goose up

CREATE TABLE USERS(
    ID UUID PRIMARY KEY,
    CREATED_AT TIMESTAMP NOT NULL,
    UPDATED_AT TIMESTAMP NOT NULL,
    NAME TEXT NOT NULL
);
-- +goose down 
DROP TABLE USERS;

-- the up statement is used to create while the down statemnt is used to delete that same table
--basically when we run in cli- goose up the commands will run and create the table 
-- and when goose down is run the command below it run usually it is to delete the table or change done