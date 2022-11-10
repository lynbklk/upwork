package dbmysql

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/wahyudibo/golang-reverse-proxy/modules/ahrefs/internal/repository"
)

type StatusRepository struct {
	db *sqlx.DB
}

func (r *StatusRepository) GetStatusesByUser(ctx context.Context, userId int64) ([]*repository.Status, error) {
	result := make([]*repository.Status, 0)
	query := fmt.Sprintf("SELECT user_id, product_id, status FROM am_user_status WHERE user_id = %d", userId)

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var status repository.Status
		rows.Scan(&status.UserId, &status.ProductId, &status.Status)
		result = append(result, &status)
	}
	return result, nil
}

func (r *StatusRepository) GetStatusesByUserAndProduct(ctx context.Context, userId int64, productIds []int64) ([]*repository.Status, error) {
	result := make([]*repository.Status, 0)

	strIds := ""
	for _, id := range productIds {
		strIds = fmt.Sprintf("%s%d,", strIds, id)
	}
	if len(strIds) != 0 {
		strIds = strIds[:len(strIds)-1]
	}

	query := fmt.Sprintf("SELECT user_id, product_id, status FROM am_user_status WHERE user_id = %d AND product_id in (%s)", userId, strIds)

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var status repository.Status
		rows.Scan(&status.UserId, &status.ProductId, &status.Status)
		result = append(result, &status)
	}
	return result, nil
}
