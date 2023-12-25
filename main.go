package main

import (
	"anteraja/backend/middleware"
	"anteraja/backend/modules/user"
	userv2 "anteraja/backend/modules/userV2"
	"anteraja/backend/utils/db"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	var (
		err error
	)
	router.Use(
		middleware.AllowCORS(),
	)

	evoDB := db.GormPostgres("host=localhost user=postgres password=Lumbanpaung,050490 dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Jakarta")
	user := user.NewRequestHandler(evoDB)
	user.Handle(router)

	userV2 := userv2.NewUserRequestHandler(evoDB)
	userV2.HandleUserV2(router)

	err = router.Run()

	if err != nil {
		log.Println("main router.Run:", err)
		return
	}

}
