package main

import (
	"log"
	"log/slog"
	"os"

	"github.com/YoungOak/GoAPI/internal/data"
	"github.com/YoungOak/GoAPI/internal/server"
)

var (
	CarManager data.Manager
	Router     server.Router

	addr string = ":8080"
)

/*
initLogger will setup the application logger and replace the
default logger with it so slog and log calls will use it.
*/
func initLogger() {
	slog.SetDefault(slog.New(
		slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{}),
	))
}

func main() {

	initLogger()

	CarManager = data.NewManager()
	Router = server.NewRouter(addr)

	Router.AddHandler("/cars", carsHandler) // GET
	Router.AddHandler("/car", carHandler)   // POST && GET && PUT

	if err := Router.Serve(); err != nil {
		log.Fatalf("Server failed during execution: %v", err)
	}
}
