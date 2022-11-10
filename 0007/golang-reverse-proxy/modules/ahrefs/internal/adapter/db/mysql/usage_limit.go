package dbmysql

import (
	"context"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"

	"github.com/wahyudibo/golang-reverse-proxy/modules/ahrefs/internal/repository"
)

type UsageLimitRepository struct {
	db *sqlx.DB
}

func (r *UsageLimitRepository) Create(ctx context.Context, userID int64) error {
	query := fmt.Sprintf("INSERT INTO ahref_usage_limit VALUES (%d, 0, 0 NOW())", userID)
	if _, err := r.db.Exec(query); err != nil {
		return err
	}

	return nil
}

func (r *UsageLimitRepository) Retrieve(ctx context.Context, userID int64) (*repository.UsageLimit, error) {
	var m repository.UsageLimit
	query := fmt.Sprintf("SELECT user_id, report_usage, export_usage, TIMESTAMPDIFF(MINUTE,last_accessed_at, NOW()) AS limit_reset_at FROM ahref_usage_limit WHERE user_id = %d", userID)

	row, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	if row.Next() {
		row.Scan(&m.UserID, &m.ReportUsage, &m.ExportUsage, &m.LimitResetAt)
	} else {
		return nil, nil
	}

	return &m, nil
}

func (r *UsageLimitRepository) Update(ctx context.Context, userID int64, reportUsage, exportUsage int, updateLastAccessedAt bool) error {
	query := fmt.Sprintf("UPDATE ahref_usage_limit SET report_usage = %d, export_usage = %d", reportUsage, exportUsage)

	if updateLastAccessedAt {
		query += ", last_accessed_at = NOW() "
	}

	query = fmt.Sprintf("%s WHERE user_id = %d", query, userID)

	res, err := r.db.Exec(query)
	if err != nil {
		return err
	}
	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return errors.New("row is not found")
	}

	return nil
}
