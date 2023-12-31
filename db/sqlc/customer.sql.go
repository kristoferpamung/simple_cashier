// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.0
// source: customer.sql

package db

import (
	"context"
)

const createCustomer = `-- name: CreateCustomer :one
INSERT INTO customers (
  customer_name,
  address
) VALUES (
  $1, $2
) RETURNING id, customer_name, address, created_at
`

type CreateCustomerParams struct {
	CustomerName string `json:"customer_name"`
	Address      string `json:"address"`
}

func (q *Queries) CreateCustomer(ctx context.Context, arg CreateCustomerParams) (Customer, error) {
	row := q.db.QueryRowContext(ctx, createCustomer, arg.CustomerName, arg.Address)
	var i Customer
	err := row.Scan(
		&i.ID,
		&i.CustomerName,
		&i.Address,
		&i.CreatedAt,
	)
	return i, err
}

const listCustomers = `-- name: ListCustomers :many
SELECT id, customer_name, address, created_at FROM customers
ORDER BY id
`

func (q *Queries) ListCustomers(ctx context.Context) ([]Customer, error) {
	rows, err := q.db.QueryContext(ctx, listCustomers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Customer{}
	for rows.Next() {
		var i Customer
		if err := rows.Scan(
			&i.ID,
			&i.CustomerName,
			&i.Address,
			&i.CreatedAt,
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
