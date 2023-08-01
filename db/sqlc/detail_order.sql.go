// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.0
// source: detail_order.sql

package db

import (
	"context"
)

const createOrderDetail = `-- name: CreateOrderDetail :one
INSERT INTO order_details (
  product_id,
  order_id,
  price,
  quantity,
  total
) VALUES (
  $1, $2, $3, $4, $5
) RETURNING id, order_id, product_id, price, quantity, total
`

type CreateOrderDetailParams struct {
	ProductID int32 `json:"product_id"`
	OrderID   int32 `json:"order_id"`
	Price     int32 `json:"price"`
	Quantity  int32 `json:"quantity"`
	Total     int32 `json:"total"`
}

func (q *Queries) CreateOrderDetail(ctx context.Context, arg CreateOrderDetailParams) (OrderDetail, error) {
	row := q.db.QueryRowContext(ctx, createOrderDetail,
		arg.ProductID,
		arg.OrderID,
		arg.Price,
		arg.Quantity,
		arg.Total,
	)
	var i OrderDetail
	err := row.Scan(
		&i.ID,
		&i.OrderID,
		&i.ProductID,
		&i.Price,
		&i.Quantity,
		&i.Total,
	)
	return i, err
}

const listOrderDetails = `-- name: ListOrderDetails :many
SELECT id, order_id, product_id, price, quantity, total FROM order_details
WHERE order_id = $1
ORDER BY id
`

func (q *Queries) ListOrderDetails(ctx context.Context, orderID int32) ([]OrderDetail, error) {
	rows, err := q.db.QueryContext(ctx, listOrderDetails, orderID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []OrderDetail{}
	for rows.Next() {
		var i OrderDetail
		if err := rows.Scan(
			&i.ID,
			&i.OrderID,
			&i.ProductID,
			&i.Price,
			&i.Quantity,
			&i.Total,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}