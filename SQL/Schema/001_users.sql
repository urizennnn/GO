-- +goose Up

    CREATE TABLE users (
        id UUID PRIMARY KEY,
        created_at TIMESTAMP NOT NULL,
        updated_at TIMESTAMP NOT NULL,
        name text not null
    );


-- +goose Down

drop tables users;