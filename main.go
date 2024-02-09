package main

import (
	"anteraja/backend/middleware"
	deposit "anteraja/backend/modules/deposite"
	"anteraja/backend/modules/role"
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

	db := db.GormPostgres("host=localhost user=postgres password=Lumbanpaung,050490 dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Jakarta")

	userV2 := userv2.NewUserRequestHandler(db)
	userV2.HandleUserV2(router)

	role := role.NewRoleRequestHandler(db)
	role.HandleRole(router)

	deposit := deposit.NewRequestDepositHandler(db)
	deposit.HandleDeposit(router)

	err = router.Run()

	if err != nil {
		log.Println("main router.Run:", err)
		return
	}

}
