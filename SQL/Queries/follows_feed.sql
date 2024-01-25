-- name: CreateFeedFollows :one

insert into follows_feed  (id,created_at,updated_at,user_id,feed_id)
values ($1,$2,$3,$4,$5)
returning *;

