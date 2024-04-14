package db

import (
	"context"
	"testing"
	"time"

	"github.com/akmshasan/fruit-store/util"
	"github.com/stretchr/testify/require"
)

func CreateRandomFruit(t *testing.T) Fruit {
	arg := CreateFruitParams{
		Name:     util.RandomName(),
		Color:    util.RandomColor(),
		Price:    int64(util.RandomPrice()),
		Quantity: int64(util.RandomQuantity()),
	}

	fruit, err := testStore.CreateFruit(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, fruit)

	require.Equal(t, arg.Name, fruit.Name)
	require.Equal(t, arg.Color, fruit.Color)
	require.Equal(t, arg.Price, fruit.Price)
	require.Equal(t, arg.Quantity, fruit.Quantity)

	require.NotZero(t, fruit.ID)
	require.NotZero(t, fruit.CreatedAt)

	return fruit
}

func TestCreateFruit(t *testing.T) {
	CreateRandomFruit(t)
}

func TestGetFruit(t *testing.T) {
	fruit1 := CreateRandomFruit(t)

	fruit2, err := testStore.GetFruit(context.Background(), fruit1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, fruit2)

	require.Equal(t, fruit1.ID, fruit2.ID)
	require.Equal(t, fruit1.Name, fruit2.Name)
	require.Equal(t, fruit1.Color, fruit2.Color)
	require.Equal(t, fruit1.Price, fruit2.Price)
	require.Equal(t, fruit1.Quantity, fruit2.Quantity)
	require.WithinDuration(t, fruit1.CreatedAt, fruit2.CreatedAt, time.Second)

}

func TestListFruits(t *testing.T) {
	arg := ListFruitsParams{
		Limit:  5,
		Offset: 5,
	}

	for i := 0; i < 10; i++ {
		CreateRandomFruit(t)
	}

	fruits, err := testStore.ListFruits(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, fruits)

	for _, fruit := range fruits {
		require.NotEmpty(t, fruit)
		require.Len(t, fruits, int(arg.Limit))
	}
}

func TestUpdateFruit(t *testing.T) {
	fruit1 := CreateRandomFruit(t)

	arg := UpdateFruitParams{
		ID:    fruit1.ID,
		Price: int64(util.RandomPrice()),
	}

	fruit2, err := testStore.UpdateFruit(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, fruit2)

	require.Equal(t, arg.ID, fruit2.ID)
	require.Equal(t, fruit1.Name, fruit2.Name)
	require.Equal(t, fruit1.Color, fruit2.Color)
	require.Equal(t, arg.Price, fruit2.Price)
	require.Equal(t, fruit1.Quantity, fruit2.Quantity)
	require.WithinDuration(t, fruit1.CreatedAt, fruit2.CreatedAt, time.Second)

}

func TestDeleteFruit(t *testing.T) {
	fruit := CreateRandomFruit(t)

	err := testStore.DeleteFruit(context.Background(), fruit.ID)
	require.NoError(t, err)
}
