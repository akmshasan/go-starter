package api

import (
	"context"
	"database/sql"
	"net/http"

	db "github.com/akmshasan/fruit-store/db/sqlc"
	"github.com/gin-gonic/gin"
)

func (server *Server) IndexPage(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "Index Page"})
}

func (server *Server) HealthStatus(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "OK"})
}

type createFruitRequest struct {
	Name     string `json:"name" binding:"required"`
	Color    string `json:"color" binding:"required"`
	Price    int64  `json:"price" binding:"required"`
	Quantity int64  `json:"quantity" binding:"required"`
}

func (server *Server) createFruit(ctx *gin.Context) {
	var req createFruitRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateFruitParams{
		Name:     req.Name,
		Color:    req.Color,
		Price:    req.Price,
		Quantity: req.Quantity,
	}

	fruit, err := server.store.CreateFruit(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusCreated, fruit)

}

type getFruitRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getFruit(ctx *gin.Context) {
	var req getFruitRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	fruit, err := server.store.GetFruit(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, fruit)
}

type listFruitRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func (server *Server) listFruit(ctx *gin.Context) {
	var req listFruitRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListFruitsParams{
		Limit:  int64(req.PageSize),
		Offset: (int64(req.PageID) - 1) * int64(req.PageSize),
	}

	fruits, err := server.store.ListFruits(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, fruits)
}

type deleteFruitRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) deleteFruit(ctx *gin.Context) {
	var req deleteFruitRequest

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	err := server.store.DeleteFruit(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusAccepted, err)
}

type updateFruitRequest struct {
	ID    int64 `form:"id" binding:"required"`
	Price int64 `form:"price" binding:"required"`
}

func (server *Server) updateFruit(ctx *gin.Context) {
	var req updateFruitRequest

	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.UpdateFruitParams{
		ID:    req.ID,
		Price: req.Price,
	}

	fruit, err := server.store.UpdateFruit(context.Background(), arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusAccepted, fruit)
}
