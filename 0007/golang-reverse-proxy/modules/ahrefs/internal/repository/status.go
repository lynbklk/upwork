package repository

import "context"

type Status struct {
	UserId    int64
	ProductId int64
	Status    int
}

type StatusRepository interface {
	GetStatusesByUser(ctx context.Context, userId int64) ([]*Status, error)
	GetStatusesByUserAndProduct(ctx context.Context, userId int64, productIds []int64) ([]*Status, error)
}
