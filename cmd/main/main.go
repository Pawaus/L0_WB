package main

import (
	"fmt"
	"l0_wb/internal/app/api"
	"l0_wb/internal/app/nats"
	"l0_wb/internal/config"
	"l0_wb/internal/database"
	"l0_wb/internal/pkg/cache"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kelseyhightower/envconfig"
)

func main() {
	var conf config.ConfigDB
	err := envconfig.Process("app", &conf)
	if err != nil {
		fmt.Println("Failed to load enviroment config DB")
	}
	var confNats config.ConfigNats
	err = envconfig.Process("nats", &confNats)
	if err != nil {
		fmt.Println("Failed to load enviroment config NATS")
	}
	fmt.Println("Load config from env")
	var db *database.Database
	for countConn := 0; countConn < 10; countConn++ {
		db = database.NewDatabase(conf.DbHost, conf.DbPort, conf.DbUser, conf.DbPassword, conf.DbName)
		if db == nil {
			fmt.Println("DB not avaliable, reconnecting ...")
			time.Sleep(3 * time.Second)
		} else {
			break
		}
	}

	if db == nil {
		fmt.Print("Failed connect to DB")
		return
	}
	db.CreateTable()
	cache := cache.New()
	nats.Setup(db, cache, confNats.ClusterName, confNats.ClientName, confNats.ServerNats)
	handler_api := api.Setup(db, cache)
	gin_server := gin.Default()
	gin_server.GET("/status", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})
	gin_server.GET("/uid/:uid", handler_api.GetUid)
	gin_server.Run()

}
