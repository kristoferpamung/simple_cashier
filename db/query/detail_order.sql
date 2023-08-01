-- name: CreateOrderDetail :one
INSERT INTO order_details (
  product_id,
  order_id,
  price,
  quantity,
  total
) VALUES (
  $1, $2, $3, $4, $5
) RETURNING *;

-- name: ListOrderDetails :many
SELECT * FROM order_details
WHERE order_id = $1
ORDER BY id;