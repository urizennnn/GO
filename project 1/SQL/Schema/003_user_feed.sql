-- +goose Up

    CREATE TABLE feed (
        id UUID PRIMARY KEY,
        created_at TIMESTAMP NOT NULL,
        updated_at TIMESTAMP NOT NULL,
        name text not null,
        url text unique not null,
        user_id UUID not null references users(id) on DELETE cascade
    );


-- +goose Down

drop table feed;