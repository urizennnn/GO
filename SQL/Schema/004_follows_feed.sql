-- +goose Up

    CREATE TABLE follows_feed (
        id UUID PRIMARY KEY,
        created_at TIMESTAMP NOT NULL,
        updated_at TIMESTAMP NOT NULL,
        user_id UUID not null references users(id) on DELETE cascade,
        feed_id UUID not null references feed(id) on DELETE cascade,
        unique(user_id,feed_id)
    );


-- +goose Down

drop table follows_feed;