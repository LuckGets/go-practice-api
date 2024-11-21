package main

import (
	"log"
	"net/http"

	"github.com/LuckGets/go-echo-rest/api/route"
	"github.com/LuckGets/go-echo-rest/boostrap"
	"github.com/LuckGets/go-echo-rest/database"
	"github.com/LuckGets/go-echo-rest/store"
	"github.com/labstack/echo/v4"
)

func main() {
	// Get the environment variable
	app := boostrap.LoadAppConfig()

	// connect to the DB
	db, err := database.NewPostgresDB(app.Env.DBurl+app.Env.DBName, app.Env.DBMaxOpenConn, app.Env.DBMaxIdleConn, app.Env.DBMaxIdleTime)

	if err != nil {
		log.Fatalf("Can't connect to the DB server:: %s", err)
	}

	// Passing the database to store
	store := store.NewStorage(db)

	e := echo.New()

	// Mount the Global middleware
	route.MountMiddleWare(e)
	// Mount the route handler
	route.MountRoute(e)

	if err := e.Start(":" + app.Env.ServerPort); err != nil && err != http.ErrServerClosed {
		log.Fatalf("The server is shutting down due to this error:: %v", err)
	}

}
