package middlewarex

import (
	"net/http"
	"sync"

	"github.com/zeromicro/go-zero/core/logx"
)

type OperationMiddleware struct {
	sync.RWMutex
	ApiPolicy map[string]bool
}

func NewOperationMiddleware() *OperationMiddleware {
	return &OperationMiddleware{
		ApiPolicy: make(map[string]bool),
	}
}

func (m *OperationMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logx.Infof("OperationMiddleware Handle")
		if m.isNeedLog(r.URL.Path) {
			logx.Infof("OperationMiddleware Handle isNeedLog")

		}

		next(w, r)
	}
}

func (m *OperationMiddleware) LoadPolicy(pl map[string]bool) error {
	m.Lock()
	defer m.Unlock()

	m.ApiPolicy = pl
	return nil
}

func (m *OperationMiddleware) isNeedLog(api string) bool {
	m.RLock()
	defer m.RUnlock()

	if v, ok := m.ApiPolicy[api]; ok {
		return v
	}

	return false
}
