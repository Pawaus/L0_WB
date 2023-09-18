package api

import (
	"l0_wb/internal/app/api/handlers"
	"l0_wb/internal/app/api/repo"
	"l0_wb/internal/app/api/usecase"
	"l0_wb/internal/database"
	"l0_wb/internal/pkg/cache"
)

func Setup(db *database.Database, cache *cache.Cache) *handlers.Handlers {
	r := repo.NewRepo(db, cache)
	u := usecase.NewUsecase(r)
	h := handlers.NewHandler(u)
	return h
}
