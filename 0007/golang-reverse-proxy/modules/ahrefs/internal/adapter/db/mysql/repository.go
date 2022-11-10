package dbmysql

import (
	"github.com/jmoiron/sqlx"

	"github.com/wahyudibo/golang-reverse-proxy/modules/ahrefs/internal/repository"
)

type Repository struct {
	usageLimit *UsageLimitRepository
	session    *SessionRepository
	status     *StatusRepository
}

func New(db *sqlx.DB) *Repository {
	usageLimitRepository := &UsageLimitRepository{
		db: db,
	}
	sessionRepository := &SessionRepository{
		db: db,
	}
	statusRepository := &StatusRepository{
		db: db,
	}
	return &Repository{
		usageLimit: usageLimitRepository,
		session:    sessionRepository,
		status:     statusRepository,
	}
}

func (r *Repository) UsageLimit() repository.UsageLimitRepository {
	return r.usageLimit
}

func (r *Repository) Session() repository.SessionRepository {
	return r.session
}

func (r *Repository) Status() repository.StatusRepository {
	return r.status
}
