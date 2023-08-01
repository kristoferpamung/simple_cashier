-- name: CreateUser :one
INSERT INTO users (
  username,
  hashed_password,
  full_name
) VALUES (
  $1, $2, $3
) RETURNING *;

-- name: ListUsers :many
SELECT * FROM users
WHERE username != "admin"
ORDER BY username;

-- name: UpdateUser :one
UPDATE users
SET photo = $1
WHERE username = $2
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users
WHERE username = $1;