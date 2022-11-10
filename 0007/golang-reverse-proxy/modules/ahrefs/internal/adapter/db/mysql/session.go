package dbmysql

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type SessionRepository struct {
	db *sqlx.DB
}

func (r *SessionRepository) FindUserIDBySession(ctx context.Context, sessionID string) (int64, error) {
	var userID int64

	query := fmt.Sprintf("SELECT user_id FROM am_session WHERE id = '%s'", sessionID)

	rows, err := r.db.Query(query)
	if err != nil {
		return 0, err
	}
	defer rows.Close()

	if rows.Next() {
		rows.Scan(&userID)
	}

	return userID, nil
}
