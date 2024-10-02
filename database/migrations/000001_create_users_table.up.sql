CREATE TABLE users (
    id serial NOT NULL PRIMARY KEY,
    tid VARCHAR(255),
    name VARCHAR(255),
    role VARCHAR(255),
    password VARCHAR(255),
    birthday TIMESTAMP
);

-- migrate -database "postgres://kristofer:mysecret@127.0.0.1:5432/joglo_dev?sslmode=false" -path database/migrations up
-- migrate create -ext sql -dir database/migrations -seq create_books_table