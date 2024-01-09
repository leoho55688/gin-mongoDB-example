package main

import (
	"time"

	route "backend/api/route"
	"backend/bootstrap"

	"github.com/gin-gonic/gin"
)

func main() {
    
    app := bootstrap.App()

    env := app.Env

    db := app.Mongo.Database(env.DBName)
    defer app.CloseDBConnection()

    timout := time.Duration(env.ContextTimeout) * time.Second

    gin := gin.Default()
    gin.SetTrustedProxies([]string{"127.0.0.1"})

    route.Setup(env, timout, db, gin)

    gin.Run((env.ServerAddress))

}