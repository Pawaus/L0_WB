package repo

import (
	"encoding/json"
	"l0_wb/internal/database"
	"l0_wb/internal/models"
	"l0_wb/internal/pkg/cache"
	"l0_wb/internal/pkg/domain"
)

type Repo struct {
	db    *database.Database
	cache *cache.Cache
}

func NewRepo(db *database.Database, cache *cache.Cache) *Repo {

	return &Repo{db: db, cache: cache}
}

func (r *Repo) ProcessOrder(uid, data string) {
	r.db.InsertData(uid, data)
	js := domain.Order{}
	json.Unmarshal([]byte(data), &js)
	r.cache.Update(js.Order_uid, js)
}

func (r *Repo) LoadCache() {
	var orders []models.Orders
	r.db.Find(&orders)
	for _, order := range orders {
		js := domain.Order{}
		json.Unmarshal(order.Data, &js)
		r.cache.Update(order.Uid, js)
	}
}
