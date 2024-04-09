// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package db

import (
	"context"
)

type Querier interface {
	CreateFruit(ctx context.Context, arg CreateFruitParams) (Fruit, error)
	DeleteFruit(ctx context.Context, id int64) error
	GetFruit(ctx context.Context, id int64) (Fruit, error)
	ListFruits(ctx context.Context, arg ListFruitsParams) ([]Fruit, error)
	UpdateFruit(ctx context.Context, arg UpdateFruitParams) (Fruit, error)
}

var _ Querier = (*Queries)(nil)