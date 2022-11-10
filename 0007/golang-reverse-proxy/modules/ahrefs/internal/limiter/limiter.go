package limiter

import (
	"context"
	"fmt"
	"github.com/wahyudibo/golang-reverse-proxy/modules/ahrefs/internal/repository"
	"net/http"
	"strings"
)

var (
	DefaultAccessLimitChecker func(next http.Handler) http.Handler
	L                         *Limiter
	productIds                = []int64{101, 102, 103, 5, 15}
)
var (
	LimitReachedTypeReport LimitReachedType = 1
	LimitReachedTypeExport LimitReachedType = 2
)

var (
	ReportUsageLimit   = 30
	ExportUsageLimit   = 10
	UsageLimitDuration = 12 * 60
)

type Limiter struct {
	ctx  context.Context
	repo repository.Repository
}

type LimitReachedType int

type LimitUsageReach struct {
	LimitReached     bool
	LimitReachedType LimitReachedType
}

func NewLimiter(repo repository.Repository) *Limiter {
	return &Limiter{
		repo: repo,
	}
}

func AccessLimitChecker(next http.Handler) http.Handler {
	return DefaultAccessLimitChecker(next)
}

func (l *Limiter) UpdateUserLimitUsage(userId int64, req *http.Request) (*LimitUsageReach, error) {
	url, method := req.URL.String(), req.Method

	// todo
	// key := fmt.Sprintf("ahx:%d:%s", userId, url)
	// is key exist in redis

	isReportConsumed, isExportConsumed := false, false
	reportUsage, exportUsage := 0, 0
	for _, uri := range ReportUsageRemarks {
		if strings.Contains(url, uri) {
			isReportConsumed = true
			break
		}
	}

	if isReportConsumed {
		for _, uri := range ReportUsageExceptRemarks {
			if strings.Contains(url, uri) {
				isReportConsumed = false
				break
			}
		}
	}

	if method == http.MethodPost {
		for _, uri := range ExportUsageRemarks {
			if strings.Contains(url, uri) {
				isExportConsumed = true
				break
			}
		}
	}

	// todo
	// regex match

	if isReportConsumed {
		reportUsage = 1
	}

	if isExportConsumed {
		exportUsage = 1
	}

	// get user's limit usage
	limitUsage, err := l.repo.UsageLimit().Retrieve(l.ctx, userId)
	if err != nil {
		return nil, err
	}

	// the user might not be registered in the ahref_usage_limit yet. We should insert data with new limit setup.
	if limitUsage == nil {
		err := l.repo.UsageLimit().Create(l.ctx, userId)
		if err != nil {
			return nil, err
		}
	}

	if limitUsage.LimitResetAt > UsageLimitDuration {
		l.repo.UsageLimit().Update(l.ctx, userId, 0, 0, true)
	}

	if !isExportConsumed && !isReportConsumed {
		return nil, nil
	}

	// get it again
	limitUsage, _ = l.repo.UsageLimit().Retrieve(l.ctx, userId)
	if limitUsage == nil {
		return nil, fmt.Errorf("some bad thing happened")
	}

	reportUsage += limitUsage.ReportUsage
	exportUsage += limitUsage.ExportUsage

	if isReportConsumed && reportUsage > ReportUsageLimit {
		return &LimitUsageReach{
			LimitReached:     true,
			LimitReachedType: LimitReachedTypeReport,
		}, nil
	}

	// todo
	// set keys in redis

	if isExportConsumed && exportUsage > ExportUsageLimit {
		return &LimitUsageReach{
			LimitReached:     true,
			LimitReachedType: LimitReachedTypeExport,
		}, nil
	}

	l.repo.UsageLimit().Update(l.ctx, userId, reportUsage, exportUsage, false)
	return nil, nil
}

func newAccessLimitChecker() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			if !strings.HasPrefix(r.RequestURI, "/ahx-static/") && !strings.HasPrefix(r.RequestURI, "/usage-limit") {
				cookie, err := r.Cookie("PHPSESSID")
				if err != nil {
					w.WriteHeader(http.StatusForbidden)
					w.Write([]byte("Access forbidden"))
					return
				}
				userId, err := L.repo.Session().FindUserIDBySession(L.ctx, cookie.Value)
				if err != nil || userId == 0 {
					w.WriteHeader(http.StatusForbidden)
					w.Write([]byte("Access forbidden"))
					return
				}

				statuses, err := L.repo.Status().GetStatusesByUserAndProduct(L.ctx, userId, productIds)
				if err != nil || len(statuses) == 0 || statuses[0].Status == 0 {
					w.WriteHeader(http.StatusForbidden)
					w.Write([]byte("Access forbidden"))
					return
				}

				usageLimit, err := L.UpdateUserLimitUsage(userId, r)
				if err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					w.Write([]byte("Internal Error"))
				}
				if usageLimit != nil && usageLimit.LimitReached {
					w.WriteHeader(http.StatusForbidden)
					w.Write([]byte("limit reached"))
				}

				found := false
				for _, status := range statuses {
					if status.Status == 1 {
						found = true
					}
				}
				if !found {
					w.WriteHeader(http.StatusForbidden)
					w.Write([]byte("Access expired"))
					return
				}
			}
			next.ServeHTTP(w, r)
		}
		return http.HandlerFunc(fn)
	}
}

func init() {
	DefaultAccessLimitChecker = newAccessLimitChecker()
}
