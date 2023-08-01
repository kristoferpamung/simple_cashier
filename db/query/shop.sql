-- name: UpdateShop :one
UPDATE shops
SET shop_name = $1,
    address = $2,
    phone_number = $3
WHERE id = 1
RETURNING *;