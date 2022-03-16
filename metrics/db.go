package metrics

import (
	"github.com/gobuffalo/envy"
	bolt "go.etcd.io/bbolt"
)

func GetDb() (*bolt.DB, error) {
	db, err := bolt.Open(envy.Get("WNF_DB", "stats.db"), 0600, nil)
	if err != nil {
		return nil, err
	}

	return db, err
}
