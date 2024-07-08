package main

import (
	"log"
	"time"

	"api.brutecore/libs/lib_db"
	"api.brutecore/libs/lib_env"
	"api.brutecore/libs/lib_jwt"

	eg "api.brutecore/internal/engine"
	handler "api.brutecore/internal/handler"
	server "api.brutecore/internal/server"
)

func main() {
	conf, err := lib_env.New()
	if err != nil {
		log.Fatalf("Error initializing configuration: %v", err)
	}
	defer conf.Close()

	db, err := lib_db.New(&lib_db.DBConfig{
		"driver":           lib_db.SQLite,
		"connectionString": conf.Database.Path,
		"logFile":          conf.Logs.DB,
	})
	if err != nil {
		log.Fatalf("Error initializing database: %v", err)
	}
	defer db.Close()

	jwt, err := lib_jwt.New(&lib_jwt.JWTConfig{
		lib_jwt.Method:          lib_jwt.MapStrToMethod(conf.Jwt.Method),
		lib_jwt.SecretKey:       []byte(conf.Jwt.Key),
		lib_jwt.AccessDuration:  conf.Jwt.AccessTime,
		lib_jwt.RefreshDuration: conf.Jwt.RefreshTime,
	})
	if err != nil {
		log.Fatalf("Error initializing JWT: %v", err)
	}

	pull := eg.NewPulling(db, time.Second*5)
	go pull.StartListen()

	app := server.New(conf)
	app.SetMiddleware()

	hdlr := handler.New(conf, db, jwt, pull)
	app.SetHandlers(hdlr)

	err = app.Listen()
	if err != nil {
		log.Fatalf("Error listening to the server: %v", err)
	}
}
