package db

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

type OrderTxParams struct {
	Username     string        `json:"username"`
	CustomerID   int32         `json:"customer_id"`
	Cash         int32         `json:"cash"`
	OrderDetails []OrderDetail `json:"order_detail"`
}

type OrderTxResult struct {
	Order       Order         `json:"order"`
	OrderDetail []OrderDetail `json:"order_detail"`
}

type Store interface {
	OrderTx(ctx context.Context, arg OrderTxParams) (OrderTxResult, error)
	Querier
}

type SQLStore struct {
	db *sql.DB
	*Queries
}

func NewStore(db *sql.DB) Store {
	return &SQLStore{
		db:      db,
		Queries: New(db),
	}
}

func (store *SQLStore) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)

	if err != nil {
		return err
	}

	q := New(tx)

	err = fn(q)

	if err != nil {
		if rbError := tx.Rollback(); rbError != nil {
			return fmt.Errorf("tx err : %v, rb err: %v", err, rbError)
		}
		return err
	}

	return tx.Commit()
}

func (store *SQLStore) OrderTx(ctx context.Context, arg OrderTxParams) (OrderTxResult, error) {
	var result OrderTxResult

	err := store.execTx(ctx, func(q *Queries) error {
		var err error
		var totalPriceOrder int
		id := time.Now().UTC().Unix()

		for _, orderDetail := range arg.OrderDetails {
			productPrice, err := getProductPrice(ctx, q, int64(orderDetail.ProductID))
			if err != nil {
				return err
			}

			quantity := orderDetail.Quantity
			totalPriceOrder += productPrice * int(quantity)
		}

		result.Order, err = q.CreateOrder(ctx, CreateOrderParams{
			ID:         id,
			Username:   arg.Username,
			CustomerID: arg.CustomerID,
			TotalPrice: int32(totalPriceOrder),
			Cash:       arg.Cash,
			Return:     arg.Cash - int32(totalPriceOrder),
		})

		if err != nil {
			return err
		}

		for _, orderDetail := range arg.OrderDetails {
			price, err := getProductPrice(ctx, q, int64(orderDetail.ProductID))

			if err != nil {
				return err
			}

			orderDetail, err := q.CreateOrderDetail(ctx, CreateOrderDetailParams{
				ProductID: orderDetail.ProductID,
				OrderID:   int32(id),
				Price:     int32(price),
				Quantity:  orderDetail.Quantity,
				Total:     int32(price) * orderDetail.Quantity,
			})
			if err != nil {
				return err
			}
			result.OrderDetail = append(result.OrderDetail, orderDetail)
		}

		return nil
	})

	return result, err
}

func getProductPrice(ctx context.Context, q *Queries, id int64) (int, error) {
	product, err := q.GetProducts(ctx, id)

	if err != nil {
		return int(product.Price), err
	}

	return int(product.Price), nil
}
