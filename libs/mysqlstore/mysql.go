package mysqlstore

import (
	"github.com/gin-contrib/sessions"
	"github.com/srinathgs/mysqlstore"
)

type Store interface {
	sessions.Store
}

func NewStore(dsn string, table string, path string, maxAge int, keyPairs ...[]byte) (Store, error) {
	s, err := mysqlstore.NewMySQLStore(dsn, table, path, maxAge, keyPairs...)
	if err != nil {
		return nil, err
	}
	return &store{s}, nil
}

type store struct {
	*mysqlstore.MySQLStore
}

func (c *store) Options(options sessions.Options) {
	c.MySQLStore.Options = options.ToGorillaOptions()
}
