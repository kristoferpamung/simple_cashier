package api

import (
	db "cashier_api/db/sqlc"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

type createCustomerRequest struct {
	CustomerName string `json:"customer_name" binding:"required"`
	Address      string `json:"address" binding:"required"`
}

func (server *Server) createCustomer(ctx *gin.Context) {
	var req createCustomerRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateCustomerParams{
		CustomerName: req.CustomerName,
		Address:      req.Address,
	}

	result, err := server.store.CreateCustomer(ctx, arg)

	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			// log.Println(pqErr.Code.Name())
			switch pqErr.Code.Name() {
			case "unique_violation":
				ctx.JSON(http.StatusForbidden, errorResponse(err))
				return
			}
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, result)
}
