package nats

import (
	"fmt"
	"l0_wb/internal/app/nats/handlers"
	"l0_wb/internal/app/nats/repo"
	"l0_wb/internal/app/nats/usecase"
	"l0_wb/internal/database"
	"l0_wb/internal/pkg/cache"
	"l0_wb/internal/pkg/nats"
)

func Setup(db *database.Database, cache *cache.Cache, natsCluster, natsClient, natsServer string) *handlers.Handler {
	stream_conn, err := nats.Connect(natsCluster, natsClient, natsServer)
	if err != nil {
		fmt.Printf("Error connect to nats %v", err.Error())
		return nil
	}
	r := repo.NewRepo(db, cache)
	u := usecase.NewUseCase(r)
	h := handlers.NewHandler(u, stream_conn)

	return h
}
