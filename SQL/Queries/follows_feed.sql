-- name: CreateFeedFollows :one

insert into follows_feed  (id,created_at,updated_at,user_id,feed_id)
values ($1,$2,$3,$4,$5)
returning *;

-- name: GetUserFeed :many

SELECT * FROM follows_feed where user_id=$1;

-- name: UnfollowFeed :exec

DELETE FROM follows_feed where id =$1 and user_id=$2;