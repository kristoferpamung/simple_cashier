-- name: CreateProduct :one
INSERT INTO products (
  product_name,
  price
) VALUES (
  $1, $2
) RETURNING *;

-- name: GetProducts :one
SELECT * FROM products
WHERE id = $1;

-- name: UpdateProduct :one
UPDATE products
SET product_name = $1,
    price = $2,
    product_image = $3
WHERE id = $4
RETURNING *;

-- name: ListProduct :many
SELECT * FROM products
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: DeleteProduct :exec
DELETE FROM products
WHERE id = $1;