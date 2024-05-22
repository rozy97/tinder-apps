package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/rozy97/tinder-apps/user-service/config"
	"github.com/rozy97/tinder-apps/user-service/handler"
	"github.com/rozy97/tinder-apps/user-service/repository"
	"github.com/rozy97/tinder-apps/user-service/usecase"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	env := config.InitEnvironment()
	db, err := sqlx.Connect("mysql", env.MysqlURI)
	if err != nil {
		log.Fatalln(err)
	}

	defer db.Close()

	repository := repository.NewRepository(db)
	usecase := usecase.NewUsecase(repository, env)
	handler := handler.NewHandler(usecase)

	r := gin.Default()

	r.POST("/register", handler.Register)
	r.POST("/login", handler.Login)
	r.POST("activation", handler.Activation)

	err = r.Run(":" + env.AppPort) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
}
