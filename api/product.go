package api

import (
	db "cashier_api/db/sqlc"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

type createProductRequest struct {
	ProductName string `json:"product_name"`
	Price       int32  `json:"price"`
}

type createProductResponse struct {
	ProductName string `json:"product_name"`
	Price       int32  `json:"price"`
}

func (server *Server) createProduct(ctx *gin.Context) {
	var req createProductRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateProductParams{
		ProductName: req.ProductName,
		Price:       req.Price,
	}

	product, err := server.store.CreateProduct(ctx, arg)

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

	rsp := createProductResponse{
		ProductName: product.ProductName,
		Price:       product.Price,
	}

	ctx.JSON(http.StatusOK, rsp)

}
