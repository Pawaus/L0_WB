package repo

import (
	"encoding/json"
	"errors"
	"l0_wb/internal/database"
	"l0_wb/internal/pkg/cache"
)

type Repo struct {
	db    *database.Database
	cache *cache.Cache
}

func NewRepo(db *database.Database, cache *cache.Cache) *Repo {

	return &Repo{db: db, cache: cache}
}

func (r *Repo) GetOrderByID(uid string) (string, error) {
	val, err := r.cache.Get(uid)
	if err != nil {
		return r.db.GetOrderById(uid), nil
	} else {
		js, err_marsh := json.Marshal(val)
		if err_marsh == nil {
			return string(js), nil
		} else {
			return "", errors.New("Failed to parce from cache")
		}

	}
}
