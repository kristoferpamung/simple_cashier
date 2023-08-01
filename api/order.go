package api

import (
	db "cashier_api/db/sqlc"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type OrderRequest struct {
	Username     string           `json:"username" binding:"required"`
	CustomerID   int32            `json:"customer_id" binding:"required"`
	Cash         int32            `json:"cash" binding:"required"`
	OrderDetails []db.OrderDetail `json:"item" binding:"required"`
}

type OrderTxResult struct {
	Order       db.Order         `json:"order"`
	OrderDetail []db.OrderDetail `json:"order_detail"`
}

type getOrderRequest struct {
	Username string `uri:"username" binding:"required"`
}

func (server *Server) createOrder(ctx *gin.Context) {
	var req OrderRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.OrderTxParams{
		Username:     req.Username,
		CustomerID:   req.CustomerID,
		Cash:         req.Cash,
		OrderDetails: req.OrderDetails,
	}

	result, err := server.store.OrderTx(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, result)

}

func (server *Server) getOrder(ctx *gin.Context) {
	var req getOrderRequest

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	fmt.Println(req.Username)

	result, err := server.store.ListOrder(ctx, req.Username)

	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, result)

}
