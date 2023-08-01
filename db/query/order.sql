-- name: CreateOrder :one
INSERT INTO orders (
  id,
  username,
  customer_id,
  total_price,
  cash,
  return
) VALUES (
  $1, $2, $3, $4, $5, $6
) RETURNING *;


-- name: ListOrder :many
SELECT * FROM orders
WHERE username = $1
ORDER BY id;