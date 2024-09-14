-- name: CreateUser :exec
INSERT INTO users(nickname, passwd) VALUES (?, ?);

-- name: CreatePost :exec
INSERT INTO posts(content, user_id) VALUES (?, ?);

-- name: GetPosts :many
SELECT user_id, content
    from posts
    order by createdAt
    limit ?
    offset ?;