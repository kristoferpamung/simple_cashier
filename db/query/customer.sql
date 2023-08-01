-- name: CreateCustomer :one
INSERT INTO customers (
  customer_name,
  address
) VALUES (
  $1, $2
) RETURNING *;

-- name: ListCustomers :many
SELECT * FROM customers
ORDER BY id;